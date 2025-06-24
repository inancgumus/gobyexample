# Listing 10.7: Testing `Shortener.Shorten`

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/4375e28a46588280e8fc48c380d3c24011ad6dea/link) / [sqlite](https://github.com/inancgumus/gobyexample/blob/4375e28a46588280e8fc48c380d3c24011ad6dea/link/sqlite) / [shortener_test.go](https://github.com/inancgumus/gobyexample/blob/4375e28a46588280e8fc48c380d3c24011ad6dea/link/sqlite/shortener_test.go)

```go
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
```

