package sqlite

import (
	"errors"
	"testing"

	"github.com/inancgumus/gobyexample/link"
)

func TestShortenerShorten(t *testing.T) {
	t.Parallel()

	lnk := link.Link{
		Key: "foo",
		URL: "https://new.link",
	}

	// Shortens a link.
	shortener := NewShortener(DialTestDB(t))

	key, err := shortener.Shorten(t.Context(), lnk)
	if err != nil {
		t.Fatalf("got err = %v, want nil", err)
	}
	if key != "foo" {
		t.Errorf(`got key %q, want "foo"`, key)
	}

	// Disallows shortening a link with a duplicate key.
	_, err = shortener.Shorten(t.Context(), lnk)
	if !errors.Is(err, link.ErrConflict) {
		t.Fatalf("\ngot err = %v\nwant ErrConflict for duplicate key", err)
	}
}
