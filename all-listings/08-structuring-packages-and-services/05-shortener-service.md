# Listing 8.5: `Shortener` service

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/2b21a63c1ad2bc4d495aea094190d69570ec5b0b/link) / [shortener.go](https://github.com/inancgumus/gobyexample/blob/2b21a63c1ad2bc4d495aea094190d69570ec5b0b/link/shortener.go)

```go
package link

import (
	"context"
	"fmt"
)

// Shortener shortens and resolves [Link] values from an in-memory storage.
type Shortener struct {
	links map[Key]Link
}

// Shorten shortens the [Link] URL and may update the [Key] if the key
// is empty. If the link is valid, it stores it in an in-memory storage.
func (s *Shortener) Shorten(_ context.Context, lnk Link) (Key, error) {
	var err error
	if lnk.Key, err = Shorten(lnk); err != nil {
		return "", fmt.Errorf("%w: %w", err, ErrBadRequest)
	}

	// Persist the link in the in-memory storage.
	if _, ok := s.links[lnk.Key]; ok {
		return "", fmt.Errorf("saving: %w", ErrConflict)
	}
	if s.links == nil {
		s.links = map[Key]Link{}
	}
	s.links[lnk.Key] = lnk

	return lnk.Key, nil
}
```

