# Listing 8.6: `Shortener.Resolve`

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/shortener.go
+++ b/link/shortener.go
@@ -31,0 +32,18 @@
+
+// Resolve resolves the [Key] to its original [Link].
+func (s *Shortener) Resolve(_ context.Context, key Key) (Link, error) {
+	if key.Empty() {
+		return Link{}, fmt.Errorf("validating: empty key: %w", ErrBadRequest)
+	}
+	if err := key.Validate(); err != nil {
+		return Link{}, fmt.Errorf("validating: %w: %w", err, ErrBadRequest)
+	}
+
+	// Retrieve the link from the in-memory storage.
+	lnk, ok := s.links[key]
+	if !ok {
+		return Link{}, fmt.Errorf("retrieving: %w", ErrNotFound)
+	}
+
+	return lnk, nil
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/b2b2537f36a790a97df719a529128b16b5eca8a4/link) / [shortener.go](https://github.com/inancgumus/gobyexample/blob/b2b2537f36a790a97df719a529128b16b5eca8a4/link/shortener.go)

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

// Resolve resolves the [Key] to its original [Link].
func (s *Shortener) Resolve(_ context.Context, key Key) (Link, error) {
	if key.Empty() {
		return Link{}, fmt.Errorf("validating: empty key: %w", ErrBadRequest)
	}
	if err := key.Validate(); err != nil {
		return Link{}, fmt.Errorf("validating: %w: %w", err, ErrBadRequest)
	}

	// Retrieve the link from the in-memory storage.
	lnk, ok := s.links[key]
	if !ok {
		return Link{}, fmt.Errorf("retrieving: %w", ErrNotFound)
	}

	return lnk, nil
}
```

