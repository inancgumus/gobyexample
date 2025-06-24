# Listing 8.3: Validation

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/link/link.go
+++ b/link/link.go
@@ -3,12 +3,35 @@
-import "strings"
+import (
+	"errors"
+	"fmt"
+	"net/url"
+	"strings"
+)
 
 // Link represents a link.
 type Link struct {
 	// URL is the original URL of the link.
 	URL string
 	// Key is the shortened key of the URL.
 	Key Key
 	// Additional fields can be added here if needed.
 }
 
+// Validate validates the [Link].
+func (lnk Link) Validate() error {
+	if err := lnk.Key.Validate(); err != nil {
+		return fmt.Errorf("key: %w", err)
+	}
+	u, err := url.ParseRequestURI(lnk.URL)
+	if err != nil {
+		return err
+	}
+	if u.Host == "" {
+		return errors.New("empty host")
+	}
+	if u.Scheme != "http" && u.Scheme != "https" {
+		return errors.New("scheme must be http or https")
+	}
+	return nil
+}
+
 // Key is the shortened key of a [Link] URL.
@@ -21,0 +45,11 @@
+
+// Validate validates the [Key].
+func (key Key) Validate() error {
+	// We use generate (8-hex-character) by default, but allow
+	// user-defined keys up to 16 characters for convenience.
+	const maxKeyLen = 16
+	if len(key.String()) > maxKeyLen {
+		return fmt.Errorf("too long (max %d)", maxKeyLen)
+	}
+	return nil
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/748efc088331be657b467fea39e207036cd8e88f/link) / [link.go](https://github.com/inancgumus/gobyexample/blob/748efc088331be657b467fea39e207036cd8e88f/link/link.go)

```go
package link

import (
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
```

