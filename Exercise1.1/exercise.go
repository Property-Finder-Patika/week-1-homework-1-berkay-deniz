// Berkay Deniz
// Exercise 1.1
// Echo program to print all command line arguments and print also the command that invoked it

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
