# Listing 2.9: Inserting a data scheme test case

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/url/url_test.go
+++ b/url/url_test.go
@@ -50,25 +50,30 @@
 var parseTests = []struct {
 	name string
 	uri  string
 	want *URL
 }{
+	{
+		name: "with_data_scheme",
+		uri:  "data:text/plain;base64,R28gYnkgRXhhbXBsZQ==",
+		want: &URL{Scheme: "data"},
+	},
 	{
 		name: "full",
 		uri:  "https://github.com/inancgumus",
 		want: &URL{
 			Scheme: "https",
 			Host:   "github.com",
 			Path:   "inancgumus",
 		},
 	},
 	{
 		name: "without_path",
 		uri:  "https://github.com",
 		want: &URL{
 			Scheme: "https",
 			Host:   "github.com",
 			Path:   "",
 		},
 	},
 	/* many more test cases can be easily added */
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/9fb6ff31ace8def0d727391e6fcb4dd047e32894/url) / [url_test.go](https://github.com/inancgumus/gobyexample/blob/9fb6ff31ace8def0d727391e6fcb4dd047e32894/url/url_test.go)

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

func TestURLString(t *testing.T) {
	u := &URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "inancgumus",
	}

	got := u.String()
	want := "https://github.com/inancgumus"
	if got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestParseWithoutPath(t *testing.T) {
	const uri = "https://github.com"

	got, err := Parse(uri)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want <nil>", uri, err)
	}

	want := &URL{
		Scheme: "https", Host: "github.com", Path: "",
	}
	if *got != *want {
		t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", uri, got, want)
	}
}

var parseTests = []struct {
	name string
	uri  string
	want *URL
}{
	{
		name: "with_data_scheme",
		uri:  "data:text/plain;base64,R28gYnkgRXhhbXBsZQ==",
		want: &URL{Scheme: "data"},
	},
	{
		name: "full",
		uri:  "https://github.com/inancgumus",
		want: &URL{
			Scheme: "https",
			Host:   "github.com",
			Path:   "inancgumus",
		},
	},
	{
		name: "without_path",
		uri:  "https://github.com",
		want: &URL{
			Scheme: "https",
			Host:   "github.com",
			Path:   "",
		},
	},
	/* many more test cases can be easily added */
}

func TestParseTable(t *testing.T) {
	for _, tt := range parseTests {
		t.Logf("run %s", tt.name)

		got, err := Parse(tt.uri)
		if err != nil {
			t.Fatalf("Parse(%q) err = %v, want <nil>", tt.uri, err)
		}
		if *got != *tt.want {
			t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", tt.uri, got, tt.want)
		}
	}
}
```

