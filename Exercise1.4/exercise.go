// Berkay Deniz
// Exercise 1.4
// Program to print duplicated lines in the input

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	linesWithFiles := make(map[string]string) // Stores the file names that contains the specific line
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "", linesWithFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error from duplicate program: %v\n", err)
				continue
			}
			countLines(f, counts, arg, linesWithFiles)
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("line %s occurs %d times in files:\n", line, n)
				fmt.Println(linesWithFiles[line])
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileName string, lines map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		lines[line] += " " + fileName
	}
}
