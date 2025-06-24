# Listing 10.6: Adding a test database helper

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/sqlite/sqlite.go
+++ b/link/sqlite/sqlite.go
@@ -3,10 +3,11 @@
 import (
 	"context"
 	"database/sql"
 	_ "embed"
 	"fmt"
+	"testing"
 
 	_ "modernc.org/sqlite"
 )
 
 //go:embed schema.sql
@@ -28,0 +30,21 @@
+
+// DialTestDB connects to a unique in-memory SQLite database.
+func DialTestDB(tb testing.TB) *sql.DB {
+	tb.Helper()
+
+	dsn := fmt.Sprintf(
+		"file:%s?mode=memory&cache=shared",
+		tb.Name(),
+	)
+	db, err := Dial(tb.Context(), dsn)
+	if err != nil {
+		tb.Fatalf("DialTestDB: %v", err)
+	}
+	tb.Cleanup(func() {
+		if err := db.Close(); err != nil {
+			tb.Logf("DialTestDB: closing: %v", err)
+		}
+	})
+
+	return db
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/e22e93f6eb1feaa5703f29b6e04f82f5ffb99e44/link) / [sqlite](https://github.com/inancgumus/gobyexample/blob/e22e93f6eb1feaa5703f29b6e04f82f5ffb99e44/link/sqlite) / [sqlite.go](https://github.com/inancgumus/gobyexample/blob/e22e93f6eb1feaa5703f29b6e04f82f5ffb99e44/link/sqlite/sqlite.go)

```go
package sqlite

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"testing"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

// Dial connects to SQLite and applies the schema for convenience.
func Dial(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening: %w", err)
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("pinging: %w", err)
	}
	if _, err := db.ExecContext(ctx, schema); err != nil {
		return nil, fmt.Errorf("applying schema: %w", err)
	}
	return db, nil
}

// DialTestDB connects to a unique in-memory SQLite database.
func DialTestDB(tb testing.TB) *sql.DB {
	tb.Helper()

	dsn := fmt.Sprintf(
		"file:%s?mode=memory&cache=shared",
		tb.Name(),
	)
	db, err := Dial(tb.Context(), dsn)
	if err != nil {
		tb.Fatalf("DialTestDB: %v", err)
	}
	tb.Cleanup(func() {
		if err := db.Close(); err != nil {
			tb.Logf("DialTestDB: closing: %v", err)
		}
	})

	return db
}
```

