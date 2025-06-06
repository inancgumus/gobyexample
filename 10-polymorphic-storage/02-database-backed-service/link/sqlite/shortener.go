package sqlite

import (
	"context"
	"database/sql"
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

// Resolve resolves a [link.Link] by its [link.Key] from the database.
func (s *Shortener) Resolve(ctx context.Context, key link.Key) (link.Link, error) {
	if key.Empty() {
		return link.Link{}, fmt.Errorf("validating: empty key: %w", link.ErrBadRequest)
	}
	if err := key.Validate(); err != nil {
		return link.Link{}, fmt.Errorf("validating: %w: %w", err, link.ErrBadRequest)
	}

	// Retrieve the link from the database.
	var uri string
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

	return link.Link{Key: key, URL: uri}, nil
}
