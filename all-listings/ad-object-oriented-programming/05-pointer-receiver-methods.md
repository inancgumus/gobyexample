# Listing D.5: Pointer receiver methods

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/47f1835361cb2579385b06fe51b3d20f78fa4f93/oop) / [receivers](https://github.com/inancgumus/gobyexample/blob/47f1835361cb2579385b06fe51b3d20f78fa4f93/oop/receivers) / [ptr](https://github.com/inancgumus/gobyexample/blob/47f1835361cb2579385b06fe51b3d20f78fa4f93/oop/receivers/ptr) / [server.go](https://github.com/inancgumus/gobyexample/blob/47f1835361cb2579385b06fe51b3d20f78fa4f93/oop/receivers/ptr/server.go)

```go
package main

import "time"

type server struct {
	url          string
	responseTime time.Duration
}

func (s *server) check() {
	s.responseTime = 3 * time.Second
}

func (s *server) slow() bool {
	return s.responseTime > 2*time.Second
}
```

