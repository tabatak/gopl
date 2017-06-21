// Echo prints the index and value of each of its command-line arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, s := range os.Args[:] {
		fmt.Printf("%d : %s\n", i, s)
	}
}
