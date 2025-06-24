# Listing 3.7: Reassembling a URL

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/url/url.go
+++ b/url/url.go
@@ -3,7 +3,6 @@
 import (
 	"errors"
-	"fmt"
 	"strings"
 )
 
 // A URL represents a parsed URL.
@@ -35,7 +34,19 @@
 // String reassembles the URL into a URL string.
 func (u *URL) String() string {
 	if u == nil {
 		return ""
 	}
-	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
+	var s string
+	if sc := u.Scheme; sc != "" {
+		s += sc
+		s += "://"
+	}
+	if h := u.Host; h != "" {
+		s += h
+	}
+	if p := u.Path; p != "" {
+		s += "/"
+		s += p
+	}
+	return s
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/9bb29680c702eb57d70d310b2c4dab9ac6cae98e/url) / [url.go](https://github.com/inancgumus/gobyexample/blob/9bb29680c702eb57d70d310b2c4dab9ac6cae98e/url/url.go)

```go
package url

import (
	"errors"
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
	var s string
	if sc := u.Scheme; sc != "" {
		s += sc
		s += "://"
	}
	if h := u.Host; h != "" {
		s += h
	}
	if p := u.Path; p != "" {
		s += "/"
		s += p
	}
	return s
}
```

