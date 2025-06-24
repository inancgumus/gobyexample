# Listing 2.11: Fixing the `Parse` function for the data scheme

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/url/url.go
+++ b/url/url.go
@@ -16,17 +16,20 @@
 // Parse parses a URL string into a URL structure.
 func Parse(rawURL string) (*URL, error) {
-	scheme, rest, ok := strings.Cut(rawURL, "://")
+	scheme, rest, ok := strings.Cut(rawURL, ":")
 	if !ok {
 		return nil, errors.New("missing scheme")
 	}
+	if !strings.HasPrefix(rest, "//") {
+		return &URL{Scheme: scheme}, nil
+	}
 
-	host, path, _ := strings.Cut(rest, "/")
+	host, path, _ := strings.Cut(rest[2:], "/")
 
 	return &URL{
 		Scheme: scheme,
 		Host:   host,
 		Path:   path,
 	}, nil
 }
 
 // String reassembles the URL into a URL string.

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/26062bac1ff0c943ef6b3419c5aa089ca8eac935/url) / [url.go](https://github.com/inancgumus/gobyexample/blob/26062bac1ff0c943ef6b3419c5aa089ca8eac935/url/url.go)

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
	if !ok {
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
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}
```

