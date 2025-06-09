# Listing E.4: A safer WaitGroup

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [concurrency](https://github.com/inancgumus/gobyexample/blob/f354f8e56da378cd4fe7b88a4ebe6d9feba31bbc/concurrency) / [syncx](https://github.com/inancgumus/gobyexample/blob/f354f8e56da378cd4fe7b88a4ebe6d9feba31bbc/concurrency/syncx) / [syncx.go](https://github.com/inancgumus/gobyexample/blob/f354f8e56da378cd4fe7b88a4ebe6d9feba31bbc/concurrency/syncx/syncx.go)

```go
package syncx

import "sync"

// SafeGroup is a safe version of [sync.WaitGroup].
type SafeGroup struct{ wg sync.WaitGroup }

// Wait waits for the group of goroutines to finish.
func (sg *SafeGroup) Wait() { sg.wg.Wait() }

// Go runs the given function in a goroutine.
func (sg *SafeGroup) Go(fn func()) {
	sg.wg.Add(1)
	go func() {
		defer sg.wg.Done()
		fn()
	}()
}
```

