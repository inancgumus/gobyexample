# Listing D.10: Gluing everything together

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/bafeec8e07b65b74570dd5ef2bbf937d25b368aa/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/bafeec8e07b65b74570dd5ef2bbf937d25b368aa/oop/interfaces) / [main.go](https://github.com/inancgumus/gobyexample/blob/bafeec8e07b65b74570dd5ef2bbf937d25b368aa/oop/interfaces/main.go)

```go
package main

import "fmt"

func main() {
	srv := &server{
		url: "auth",
	}
	srv.check()
	if !srv.slow() {
		return
	}
	msg := fmt.Sprintf("%s server is slow: %s", srv.url, srv.responseTime)
	notify(new(slackNotifier), msg)
	notify(new(smsNotifier), msg)
}
```

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/bafeec8e07b65b74570dd5ef2bbf937d25b368aa/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/bafeec8e07b65b74570dd5ef2bbf937d25b368aa/oop/interfaces) / [server.go](https://github.com/inancgumus/gobyexample/blob/bafeec8e07b65b74570dd5ef2bbf937d25b368aa/oop/interfaces/server.go)

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

