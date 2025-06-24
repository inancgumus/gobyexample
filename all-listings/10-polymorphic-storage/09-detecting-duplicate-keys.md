# Listing 10.9: Detecting duplicate keys

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/sqlite/shortener.go
+++ b/link/sqlite/shortener.go
@@ -21,19 +21,22 @@
 // Shorten shortens the URL of the [link.Link] and returns a [link.Key].
 func (s *Shortener) Shorten(ctx context.Context, lnk link.Link) (link.Key, error) {
 	var err error
 	if lnk.Key, err = link.Shorten(lnk); err != nil {
 		return "", fmt.Errorf("%w: %w", err, link.ErrBadRequest)
 	}
 
 	// Persist the link in the database.
 	_, err = s.db.ExecContext(
 		ctx,
 		`INSERT INTO links (short_key, uri) VALUES ($1, $2)`,
 		lnk.Key, lnk.URL,
 	)
+	if isPrimaryKeyViolation(err) {
+		return "", fmt.Errorf("saving: %w", link.ErrConflict)
+	}
 	if err != nil {
 		return "", fmt.Errorf("saving: %w: %w", err, link.ErrInternal)
 	}
 
 	return lnk.Key, nil
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/5ca0b4187900a93f0676a5e1eca3e14dd2226282/link) / [sqlite](https://github.com/inancgumus/gobyexample/blob/5ca0b4187900a93f0676a5e1eca3e14dd2226282/link/sqlite) / [shortener.go](https://github.com/inancgumus/gobyexample/blob/5ca0b4187900a93f0676a5e1eca3e14dd2226282/link/sqlite/shortener.go)

```go
package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/inancgumus/gobyexample/link"
)

// Shortener is a link shortener service that is backed by SQLite.
type Shortener struct {
	db *sql.DB
}

// NewShortener returns a new [Shortener] service.
func NewShortener(db *sql.DB) *Shortener {
	return &Shortener{db: db}
}

// Shorten shortens the URL of the [link.Link] and returns a [link.Key].
func (s *Shortener) Shorten(ctx context.Context, lnk link.Link) (link.Key, error) {
	var err error
	if lnk.Key, err = link.Shorten(lnk); err != nil {
		return "", fmt.Errorf("%w: %w", err, link.ErrBadRequest)
	}

	// Persist the link in the database.
	_, err = s.db.ExecContext(
		ctx,
		`INSERT INTO links (short_key, uri) VALUES ($1, $2)`,
		lnk.Key, lnk.URL,
	)
	if isPrimaryKeyViolation(err) {
		return "", fmt.Errorf("saving: %w", link.ErrConflict)
	}
	if err != nil {
		return "", fmt.Errorf("saving: %w: %w", err, link.ErrInternal)
	}

	return lnk.Key, nil
}
```

