// Experiment to measure the difference in two different way of prints command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	benchmark1()
	benchmark2()
}

func benchmark1() {
	start := time.Now()

	// inefficient version
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func benchmark2() {
	start := time.Now()

	// strings.Join
	fmt.Println(strings.Join(os.Args[1:], " "))

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
