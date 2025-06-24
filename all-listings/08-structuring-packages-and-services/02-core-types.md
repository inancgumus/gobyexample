# Listing 8.2: Core types

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/63eca32966f26d4fdf6302c740d8647aeed7747b/link) / [link.go](https://github.com/inancgumus/gobyexample/blob/63eca32966f26d4fdf6302c740d8647aeed7747b/link/link.go)

```go
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
```

