# Listing 2.7: A table-driven test case

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/url/url_test.go
+++ b/url/url_test.go
@@ -48,0 +49,26 @@
+
+var parseTests = []struct {
+	name string
+	uri  string
+	want *URL
+}{
+	{
+		name: "full",
+		uri:  "https://github.com/inancgumus",
+		want: &URL{
+			Scheme: "https",
+			Host:   "github.com",
+			Path:   "inancgumus",
+		},
+	},
+	{
+		name: "without_path",
+		uri:  "https://github.com",
+		want: &URL{
+			Scheme: "https",
+			Host:   "github.com",
+			Path:   "",
+		},
+	},
+	/* many more test cases can be easily added */
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/b55f5e2245e6db495894bed932beade825fa7537/url) / [url_test.go](https://github.com/inancgumus/gobyexample/blob/b55f5e2245e6db495894bed932beade825fa7537/url/url_test.go)

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

