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
