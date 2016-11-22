package compressible_test

import (
	"fmt"

	_ "github.com/GitbookIO/mimedb/autoload"
	"github.com/go-http-utils/compressible"
)

func ExampleTest() {
	fmt.Println(compressible.Test("text/html"))
	// -> true

	fmt.Println(compressible.Test("image/png"))
	// -> false
}
