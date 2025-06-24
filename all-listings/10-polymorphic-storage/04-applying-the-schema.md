# Listing 10.4: Applying the schema

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/sqlite/sqlite.go
+++ b/link/sqlite/sqlite.go
@@ -3,19 +3,26 @@
 import (
 	"context"
 	"database/sql"
+	_ "embed"
 	"fmt"
 
 	_ "modernc.org/sqlite"
 )
 
+//go:embed schema.sql
+var schema string
+
 // Dial connects to SQLite and applies the schema for convenience.
 func Dial(ctx context.Context, dsn string) (*sql.DB, error) {
 	db, err := sql.Open("sqlite", dsn)
 	if err != nil {
 		return nil, fmt.Errorf("opening: %w", err)
 	}
 	if err := db.PingContext(ctx); err != nil {
 		return nil, fmt.Errorf("pinging: %w", err)
 	}
+	if _, err := db.ExecContext(ctx, schema); err != nil {
+		return nil, fmt.Errorf("applying schema: %w", err)
+	}
 	return db, nil
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/451413644f00cd298a9bf2a2eaeccc2310094ff7/link) / [sqlite](https://github.com/inancgumus/gobyexample/blob/451413644f00cd298a9bf2a2eaeccc2310094ff7/link/sqlite) / [sqlite.go](https://github.com/inancgumus/gobyexample/blob/451413644f00cd298a9bf2a2eaeccc2310094ff7/link/sqlite/sqlite.go)

```go
package sqlite

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

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
```

