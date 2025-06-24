# Listing D.4: Returning the mutated value

## What's changed?

> [!TIP]
> You can copy and directly [git-apply](https://tldr.inbrowser.app/pages/common/git-apply) this diff to your local copy of the code.

```diff
--- a/oop/receivers/value/usage.go
+++ b/oop/receivers/value/usage.go
@@ -7,2 +7,2 @@
-func (u usage) high() bool { return u >= 95 }
-func (u usage) set(to int) { u = usage(to) }
+func (u usage) high() bool       { return u >= 95 }
+func (u usage) set(to int) usage { return usage(to) }
@@ -10,5 +10,5 @@
 func main() {
-	var cpu usage = 99
-	cpu.set(70)
+	var cpu usage = 99 // cpu is 99
+	cpu = cpu.set(70)  // cpu is 70
 	fmt.Printf("cpu: %d%% high:%t\n", cpu, cpu.high())
 }

```
## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [oop](https://github.com/inancgumus/gobyexample/blob/bf61dccd09b7c5887c0698c7d7e54dc07ecc0e3f/oop) / [receivers](https://github.com/inancgumus/gobyexample/blob/bf61dccd09b7c5887c0698c7d7e54dc07ecc0e3f/oop/receivers) / [value](https://github.com/inancgumus/gobyexample/blob/bf61dccd09b7c5887c0698c7d7e54dc07ecc0e3f/oop/receivers/value) / [usage.go](https://github.com/inancgumus/gobyexample/blob/bf61dccd09b7c5887c0698c7d7e54dc07ecc0e3f/oop/receivers/value/usage.go)

```go
package main

import "fmt"

type usage int

func (u usage) high() bool       { return u >= 95 }
func (u usage) set(to int) usage { return usage(to) }

func main() {
	var cpu usage = 99 // cpu is 99
	cpu = cpu.set(70)  // cpu is 70
	fmt.Printf("cpu: %d%% high:%t\n", cpu, cpu.high())
}
```

