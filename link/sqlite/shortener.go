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
	if err != nil {
		return "", fmt.Errorf("saving: %w: %w", err, link.ErrInternal)
	}

	return lnk.Key, nil
}
