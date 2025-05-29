# Listing 3.5: Fix for a nil pointer

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/url/url.go
+++ b/url/url.go
@@ -35,4 +35,7 @@
 // String reassembles the URL into a URL string.
 func (u *URL) String() string {
+	if u == nil {
+		return ""
+	}
 	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/139b059ba8f38ab8cb0c9aa1e2a5e385537a4de0/url) / [url.go](https://github.com/inancgumus/gobyexample/blob/139b059ba8f38ab8cb0c9aa1e2a5e385537a4de0/url/url.go)

```go
package url

import (
	"errors"
	"fmt"
	"strings"
)

// A URL represents a parsed URL.
type URL struct {
	Scheme string
	Host   string
	Path   string
}

// Parse parses a URL string into a URL structure.
func Parse(rawURL string) (*URL, error) {
	scheme, rest, ok := strings.Cut(rawURL, ":")
	if !ok || scheme == "" {
		return nil, errors.New("missing scheme")
	}
	if !strings.HasPrefix(rest, "//") {
		return &URL{Scheme: scheme}, nil
	}

	host, path, _ := strings.Cut(rest[2:], "/")

	return &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}, nil
}

// String reassembles the URL into a URL string.
func (u *URL) String() string {
	if u == nil {
		return ""
	}
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}
```

