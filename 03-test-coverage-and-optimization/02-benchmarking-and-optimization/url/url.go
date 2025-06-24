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
