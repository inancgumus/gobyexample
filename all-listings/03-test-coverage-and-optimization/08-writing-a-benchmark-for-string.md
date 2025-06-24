# Listing 3.8: Writing a benchmark for `String`

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/url/url_test.go
+++ b/url/url_test.go
@@ -140,0 +141,11 @@
+
+func BenchmarkURLString(b *testing.B) {
+	u := &URL{
+		Scheme: "https",
+		Host:   "github.com",
+		Path:   "inancgumus",
+	}
+	for b.Loop() {
+		_ = u.String()
+	}
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [url](https://github.com/inancgumus/gobyexample/blob/def1b4c461d0e3aa2f16d8530815cea3c20ee7a5/url) / [url_test.go](https://github.com/inancgumus/gobyexample/blob/def1b4c461d0e3aa2f16d8530815cea3c20ee7a5/url/url_test.go)

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
	tests := []struct {
		name string
		uri  *URL
		want string
	}{
		{
			name: "nil",
			uri:  nil,
			want: "",
		},
		{
			name: "empty",
			uri:  new(URL),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.uri.String()
			if got != tt.want {
				t.Errorf("\ngot  %q\nwant %q\nfor  %#v", got, tt.want, tt.uri)
			}
		})
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

func TestParseSubtests(t *testing.T) {
	// We can put common test setup and teardown logic here.

	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.uri)
			if err != nil {
				t.Fatalf("Parse(%q) err = %v, want <nil>", tt.uri, err)
			}
			if *got != *tt.want {
				t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", tt.uri, got, tt.want)
			}
		})
	}
}

func TestParseError(t *testing.T) {
	tests := []struct {
		name string
		uri  string
	}{
		{name: "without_scheme", uri: "github.com"},
		{name: "empty_scheme", uri: "://github.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.uri)
			if err == nil {
				t.Errorf("Parse(%q) err=nil; want an error", tt.uri)
			}
		})
	}
}

func BenchmarkURLString(b *testing.B) {
	u := &URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "inancgumus",
	}
	for b.Loop() {
		_ = u.String()
	}
}
```

