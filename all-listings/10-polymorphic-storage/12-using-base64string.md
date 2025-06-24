# Listing 10.12: Using `base64String`

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/sqlite/shortener.go
+++ b/link/sqlite/shortener.go
@@ -24,48 +24,48 @@
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
-		lnk.Key, lnk.URL,
+		lnk.Key, base64String(lnk.URL),
 	)
 	if isPrimaryKeyViolation(err) {
 		return "", fmt.Errorf("saving: %w", link.ErrConflict)
 	}
 	if err != nil {
 		return "", fmt.Errorf("saving: %w: %w", err, link.ErrInternal)
 	}
 
 	return lnk.Key, nil
 }
 
 // Resolve resolves a [link.Link] by its [link.Key] from the database.
 func (s *Shortener) Resolve(ctx context.Context, key link.Key) (link.Link, error) {
 	if key.Empty() {
 		return link.Link{}, fmt.Errorf("validating: empty key: %w", link.ErrBadRequest)
 	}
 	if err := key.Validate(); err != nil {
 		return link.Link{}, fmt.Errorf("validating: %w: %w", err, link.ErrBadRequest)
 	}
 
 	// Retrieve the link from the database.
-	var uri string
+	var uri base64String
 	err := s.db.QueryRowContext(
 		ctx,
 		`SELECT uri FROM links WHERE short_key = $1`,
 		key,
 	).Scan(&uri)
 	if errors.Is(err, sql.ErrNoRows) {
 		return link.Link{}, link.ErrNotFound
 	}
 	if err != nil {
 		return link.Link{}, fmt.Errorf("retrieving: %w: %w", err, link.ErrInternal)
 	}
 
-	return link.Link{Key: key, URL: uri}, nil
+	return link.Link{Key: key, URL: uri.String()}, nil
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/2fcbd9374867e0566fcf65621dae7f1f86339e0b/link) / [sqlite](https://github.com/inancgumus/gobyexample/blob/2fcbd9374867e0566fcf65621dae7f1f86339e0b/link/sqlite) / [shortener.go](https://github.com/inancgumus/gobyexample/blob/2fcbd9374867e0566fcf65621dae7f1f86339e0b/link/sqlite/shortener.go)

```go
package sqlite

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/base64"
	"errors"
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
		lnk.Key, base64String(lnk.URL),
	)
	if isPrimaryKeyViolation(err) {
		return "", fmt.Errorf("saving: %w", link.ErrConflict)
	}
	if err != nil {
		return "", fmt.Errorf("saving: %w: %w", err, link.ErrInternal)
	}

	return lnk.Key, nil
}

// Resolve resolves a [link.Link] by its [link.Key] from the database.
func (s *Shortener) Resolve(ctx context.Context, key link.Key) (link.Link, error) {
	if key.Empty() {
		return link.Link{}, fmt.Errorf("validating: empty key: %w", link.ErrBadRequest)
	}
	if err := key.Validate(); err != nil {
		return link.Link{}, fmt.Errorf("validating: %w: %w", err, link.ErrBadRequest)
	}

	// Retrieve the link from the database.
	var uri base64String
	err := s.db.QueryRowContext(
		ctx,
		`SELECT uri FROM links WHERE short_key = $1`,
		key,
	).Scan(&uri)
	if errors.Is(err, sql.ErrNoRows) {
		return link.Link{}, link.ErrNotFound
	}
	if err != nil {
		return link.Link{}, fmt.Errorf("retrieving: %w: %w", err, link.ErrInternal)
	}

	return link.Link{Key: key, URL: uri.String()}, nil
}

type base64String string

// Value implements the driver.Valuer interface.
func (bs base64String) Value() (driver.Value, error) {
	return base64.StdEncoding.EncodeToString([]byte(bs)), nil
}

// Scan implements the sql.Scanner interface.
func (bs *base64String) Scan(src any) error {
	ss, ok := src.(string)
	if !ok {
		return fmt.Errorf("decoding: %q is %T, not string", ss, src)
	}
	dst, err := base64.StdEncoding.DecodeString(ss)
	if err != nil {
		return fmt.Errorf("decoding %q: %w", ss, err)
	}
	*bs = base64String(dst)
	return nil
}

// String implements the fmt.Stringer interface.
func (bs base64String) String() string {
	return string(bs)
}
```

