# Listing 4.5: Setting sensible defaults

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/hit.go
+++ b/hit/cmd/hit/hit.go
@@ -26,8 +26,11 @@
 func main() {
-	var c config
+	c := config{
+		n: 100,
+		c: 1,
+	}
 	if err := parseArgs(&c, os.Args[1:]); err != nil {
 		fmt.Printf("%s\n%s", err, usage)
 		os.Exit(1)
 	}
 	fmt.Printf("%s\n\nSending %d requests to %q (concurrency: %d)\n", logo, c.n, c.url, c.c)
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/9e0485850ba33ef6af26c3ca6e3f5bcee285aeeb/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/9e0485850ba33ef6af26c3ca6e3f5bcee285aeeb/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/9e0485850ba33ef6af26c3ca6e3f5bcee285aeeb/hit/cmd/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/9e0485850ba33ef6af26c3ca6e3f5bcee285aeeb/hit/cmd/hit/hit.go)

```go
package main

import (
	"fmt"
	"os"
)

const logo = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/`

const usage = `  
Usage:
  -url
       HTTP server URL (required)
  -n
       Number of requests
  -c
       Concurrency level
  -rps
       Requests per second`

func main() {
	c := config{
		n: 100,
		c: 1,
	}
	if err := parseArgs(&c, os.Args[1:]); err != nil {
		fmt.Printf("%s\n%s", err, usage)
		os.Exit(1)
	}
	fmt.Printf("%s\n\nSending %d requests to %q (concurrency: %d)\n", logo, c.n, c.url, c.c)
}
```

