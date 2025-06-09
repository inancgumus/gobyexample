# Listing 6.7: Testing Summarize

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/ae369b9e65d7c460e6aae89f65b0298a10163756/hit) / [result_test.go](https://github.com/inancgumus/gobyexample/blob/ae369b9e65d7c460e6aae89f65b0298a10163756/hit/result_test.go)

```go
package hit

import (
	"slices"
	"testing"
	"time"
)

func TestSummarizeFastestResult(t *testing.T) {
	results := []Result{
		{Duration: 2 * time.Second},
		{Duration: 5 * time.Second},
	}
	sum := Summarize(Results(slices.Values(results)))
	if sum.Fastest != 2*time.Second {
		t.Errorf("Fastest=%v; want 2s", sum.Fastest)
	}
}

func TestSummarizeNilResults(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("should not panic: %v", err)
		}
	}()
	_ = Summarize(nil)
}
```

