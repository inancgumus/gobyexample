package link

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// Link represents a link.
type Link struct {
	// URL is the original URL of the link.
	URL string
	// Key is the shortened key of the URL.
	Key Key
	// Additional fields can be added here if needed.
}

// Validate validates the [Link].
func (lnk Link) Validate() error {
	if err := lnk.Key.Validate(); err != nil {
		return fmt.Errorf("key: %w", err)
	}
	u, err := url.ParseRequestURI(lnk.URL)
	if err != nil {
		return err
	}
	if u.Host == "" {
		return errors.New("empty host")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("scheme must be http or https")
	}
	return nil
}

// Key is the shortened key of a [Link] URL.
type Key string

// Shorten shortens the [Link] URL and generates a new [Key]
// if the [Key] is empty. Otherwise, it returns the same [Key].
// It returns an error if the [Link] is invalid.
func Shorten(lnk Link) (Key, error) {
	if lnk.Key.Empty() {
		// We use SHA256 (6-base64 characters) by default.
		// This is a good compromise between uniqueness and length.
		// In production, you might want to use a more sophisticated
		// approach, like a database to ensure uniqueness.
		sum := sha256.Sum256([]byte(lnk.URL))
		lnk.Key = Key(base64.RawURLEncoding.EncodeToString(sum[:6]))
	}
	if err := lnk.Validate(); err != nil {
		return "", fmt.Errorf("validating: %w", err)
	}
	return lnk.Key, nil
}

// String returns the key without leading or trailing spaces.
func (key Key) String() string { return strings.TrimSpace(string(key)) }

// Empty reports whether the [Key] is empty.
func (key Key) Empty() bool { return key.String() == "" }

// Validate validates the [Key].
func (key Key) Validate() error {
	// We use generate (8-hex-character) by default, but allow
	// user-defined keys up to 16 characters for convenience.
	const maxKeyLen = 16
	if len(key.String()) > maxKeyLen {
		return fmt.Errorf("too long (max %d)", maxKeyLen)
	}
	return nil
}
