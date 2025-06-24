# Listing 2.1: Implementing the `url` package

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/1cedc61d8fbdc46ea14e67db385eb0cc31e43f3d/url) / [url.go](https://github.com/inancgumus/gobyexample/blob/1cedc61d8fbdc46ea14e67db385eb0cc31e43f3d/url/url.go)

```go
package url

import "fmt"

// A URL represents a parsed URL.
type URL struct {
	Scheme string
	Host   string
	Path   string
}

// Parse parses a URL string into a URL structure.
func Parse(rawURL string) (*URL, error) {
	return &URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "inancgumus",
	}, nil
}

// String reassembles the URL into a URL string.
func (u *URL) String() string {
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}
```

