# Listing 5.9: Testing with valid flags

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/hit/cmd/hit/config_test.go
+++ b/hit/cmd/hit/config_test.go
@@ -3,5 +3,57 @@
+import (
+	"io"
+	"testing"
+)
+
 type parseArgsTest struct {
 	name string
 	args []string
 	want config
 }
+
+func TestParseArgsValidInput(t *testing.T) {
+	t.Parallel()
+
+	for _, tt := range []parseArgsTest{
+		{
+			name: "all_flags",
+			args: []string{"-n=10", "-c=5", "-rps=5", "http://test"},
+			want: config{n: 10, c: 5, rps: 5, url: "http://test"},
+		},
+
+		// exercise: test with a mixture of flags
+	} {
+		t.Run(tt.name, func(t *testing.T) {
+			t.Parallel()
+
+			var got config
+			if err := parseArgs(&got, tt.args, io.Discard); err != nil {
+				t.Fatalf("parseArgs() error = %v, want no error", err)
+			}
+			if got != tt.want {
+				t.Errorf("flags = %+v, want %+v", got, tt.want)
+			}
+		})
+	}
+}
+
+func TestParseArgsInvalidInput(t *testing.T) {
+	t.Parallel()
+
+	for _, tt := range []parseArgsTest{
+		{name: "n_syntax", args: []string{"-n=ONE", "http://test"}},
+		{name: "n_zero", args: []string{"-n=0", "http://test"}},
+		{name: "n_negative", args: []string{"-n=-1", "http://test"}},
+
+		// exercise: test other error conditions
+	} {
+		t.Run(tt.name, func(t *testing.T) {
+			t.Parallel()
+
+			err := parseArgs(&config{}, tt.args, io.Discard)
+			if err == nil {
+				t.Fatal("parseArgs() = nil, want error")
+			}
+		})
+	}
+}

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/33d1fa5f74861ee3b8b92fd58b803087688eb96e/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/33d1fa5f74861ee3b8b92fd58b803087688eb96e/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/33d1fa5f74861ee3b8b92fd58b803087688eb96e/hit/cmd/hit) / [config_test.go](https://github.com/inancgumus/gobyexample/blob/33d1fa5f74861ee3b8b92fd58b803087688eb96e/hit/cmd/hit/config_test.go)

```go
package main

import (
	"io"
	"testing"
)

type parseArgsTest struct {
	name string
	args []string
	want config
}

func TestParseArgsValidInput(t *testing.T) {
	t.Parallel()

	for _, tt := range []parseArgsTest{
		{
			name: "all_flags",
			args: []string{"-n=10", "-c=5", "-rps=5", "http://test"},
			want: config{n: 10, c: 5, rps: 5, url: "http://test"},
		},

		// exercise: test with a mixture of flags
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got config
			if err := parseArgs(&got, tt.args, io.Discard); err != nil {
				t.Fatalf("parseArgs() error = %v, want no error", err)
			}
			if got != tt.want {
				t.Errorf("flags = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestParseArgsInvalidInput(t *testing.T) {
	t.Parallel()

	for _, tt := range []parseArgsTest{
		{name: "n_syntax", args: []string{"-n=ONE", "http://test"}},
		{name: "n_zero", args: []string{"-n=0", "http://test"}},
		{name: "n_negative", args: []string{"-n=-1", "http://test"}},

		// exercise: test other error conditions
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := parseArgs(&config{}, tt.args, io.Discard)
			if err == nil {
				t.Fatal("parseArgs() = nil, want error")
			}
		})
	}
}
```

