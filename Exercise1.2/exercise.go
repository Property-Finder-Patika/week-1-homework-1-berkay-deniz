// Berkay Deniz
// Exercise 1.2
// echo program to print the index and value of each of its arguments, one per line.

package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Print(index + 1)
		fmt.Print(" ")
		fmt.Print(arg)
		fmt.Println("")
	}
}
