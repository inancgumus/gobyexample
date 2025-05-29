# Listing 5.8: Adding a shared test case type

## Code in the file

> [!TIP]
> Click the links to see the file and its directory in their original locations and state as they were at the time of the listing.

## [hit](https://github.com/inancgumus/gobyexample/blob/75bc87152d77edc95bcf66a0fcfbd6b2109257d9/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/75bc87152d77edc95bcf66a0fcfbd6b2109257d9/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/75bc87152d77edc95bcf66a0fcfbd6b2109257d9/hit/cmd/hit) / [config_test.go](https://github.com/inancgumus/gobyexample/blob/75bc87152d77edc95bcf66a0fcfbd6b2109257d9/hit/cmd/hit/config_test.go)

```go
package main

type parseArgsTest struct {
	name string
	args []string
	want config
}
```

