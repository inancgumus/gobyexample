# Listing 3.11: Parallel tests

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/618b476596cda5d76464238ed3f4f1fec7012ecb/url) / [parallel_test.go](https://github.com/inancgumus/gobyexample/blob/618b476596cda5d76464238ed3f4f1fec7012ecb/url/parallel_test.go)

```go
package url

import (
	"testing"
	"time"
)

func TestParallelOne(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Second)
}

func TestParallelTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Second)
}

func TestSequential(t *testing.T) {
}
```

