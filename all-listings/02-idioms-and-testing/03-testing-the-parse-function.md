# Listing 2.3: Testing the `Parse` function

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/e579a9579f8e48ec6e5937fddfb27ab62ee5250b/url) / [url_test.go](https://github.com/inancgumus/gobyexample/blob/e579a9579f8e48ec6e5937fddfb27ab62ee5250b/url/url_test.go)

```go
package url

import "testing"

func TestParse(t *testing.T) {
	const uri = "https://github.com/inancgumus"

	got, err := Parse(uri)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want <nil>", uri, err)
	}
	want := &URL{
		Scheme: "https", Host: "github.com", Path: "inancgumus",
	}
	if *got != *want {
		t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", uri, got, want)
	}
}
```

