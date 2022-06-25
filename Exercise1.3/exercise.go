// Berkay Deniz
// Exercise 1.3
// Experiment for the echo function. Compares the execution time of both versions.

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// First version. Inefficient. Uses string concatenation
	exampleArgs := make([]string, 1000)

	start := time.Now()
	var s, sep string
	for i := 0; i < len(exampleArgs); i++ {
		s += sep + exampleArgs[i]
		sep = " "
	}
	fmt.Print(s)
	fmt.Printf("\n")
	elapsed := time.Since(start)
	fmt.Printf("Execution time for the first version: %s\n", elapsed)

	// Second version. Efficient. Usses strings.Join
	start2 := time.Now()
	fmt.Print(strings.Join(exampleArgs, " "))
	fmt.Printf("\n")
	elapsed2 := time.Since(start2)
	fmt.Printf("Execution time for the second version: %s\n", elapsed2)

}
