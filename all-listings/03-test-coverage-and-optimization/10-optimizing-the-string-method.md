# Listing 3.10: Optimizing the `String` method

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/url/url.go
+++ b/url/url.go
@@ -34,19 +34,31 @@
 // String reassembles the URL into a URL string.
 func (u *URL) String() string {
 	if u == nil {
 		return ""
 	}
-	var s string
+
+	const (
+		lenSchemeSeparator = len("://")
+		lenPathSeparator   = len("/")
+	)
+	lenURL := len(u.Scheme) + lenSchemeSeparator +
+		len(u.Host) + lenPathSeparator +
+		len(u.Path)
+
+	var s strings.Builder
+	s.Grow(lenURL)
+
 	if sc := u.Scheme; sc != "" {
-		s += sc
-		s += "://"
+		s.WriteString(sc)
+		s.WriteString("://")
 	}
 	if h := u.Host; h != "" {
-		s += h
+		s.WriteString(h)
 	}
 	if p := u.Path; p != "" {
-		s += "/"
-		s += p
+		s.WriteByte('/')
+		s.WriteString(p)
 	}
-	return s
+
+	return s.String()
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/f8e04e737bf63ff4d4e344ab691e8ee05e0eb292/url) / [url.go](https://github.com/inancgumus/gobyexample/blob/f8e04e737bf63ff4d4e344ab691e8ee05e0eb292/url/url.go)

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

	const (
		lenSchemeSeparator = len("://")
		lenPathSeparator   = len("/")
	)
	lenURL := len(u.Scheme) + lenSchemeSeparator +
		len(u.Host) + lenPathSeparator +
		len(u.Path)

	var s strings.Builder
	s.Grow(lenURL)

	if sc := u.Scheme; sc != "" {
		s.WriteString(sc)
		s.WriteString("://")
	}
	if h := u.Host; h != "" {
		s.WriteString(h)
	}
	if p := u.Path; p != "" {
		s.WriteByte('/')
		s.WriteString(p)
	}

	return s.String()
}
```

