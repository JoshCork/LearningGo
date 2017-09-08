// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.

package main

import "fmt"

func main() {

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 5; n++ {
		fmt.Println(n % 2)
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
