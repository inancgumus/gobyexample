# Listing 9.11: traceid middleware

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [link](https://github.com/inancgumus/gobyexample/blob/3adb7919bb212ed132ee9bd82cca339ffd2b75b5/link) / [kit](https://github.com/inancgumus/gobyexample/blob/3adb7919bb212ed132ee9bd82cca339ffd2b75b5/link/kit) / [traceid](https://github.com/inancgumus/gobyexample/blob/3adb7919bb212ed132ee9bd82cca339ffd2b75b5/link/kit/traceid) / [http.go](https://github.com/inancgumus/gobyexample/blob/3adb7919bb212ed132ee9bd82cca339ffd2b75b5/link/kit/traceid/http.go)

```go
package traceid

import "net/http"

// Middleware adds a new trace ID to [http.Request.Context].
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := FromContext(r.Context()); !ok {
			ctx := WithContext(r.Context(), New())
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}
```

