package url_test

import (
	"fmt"
	"log"

	"github.com/inancgumus/gobyexample/url"
)

func ExampleParse() {
	uri, err := url.Parse("https://github.com/inancgumus")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uri)
	// Output:
	// https://github.com/inancgumus
}
