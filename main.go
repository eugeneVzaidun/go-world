// init main package
package main

// import fmt package
import (
	"fmt"

	"github.com/pkg/errors"
)

// main function
func main() {
	// print hello world
	fmt.Println("Hello World")
	// print new error
	fmt.Println(errors.New("new error"))
	
}
