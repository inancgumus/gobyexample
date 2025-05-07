package link

import "strings"

// Link represents a link.
type Link struct {
	// URL is the original URL of the link.
	URL string
	// Key is the shortened key of the URL.
	Key Key
	// Additional fields can be added here if needed.
}

// Key is the shortened key of a [Link] URL.
type Key string

// String returns the key without leading or trailing spaces.
func (key Key) String() string { return strings.TrimSpace(string(key)) }

// Empty reports whether the [Key] is empty.
func (key Key) Empty() bool { return key.String() == "" }
