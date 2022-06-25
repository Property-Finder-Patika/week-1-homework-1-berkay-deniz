// Berkay Deniz
package main

import (
	"fmt"
	"strings"
)

func main() {
	greetPrinter(createGreetInTurkish, "Berkay")
	greetPrinter(createGreetInEnglish, "John")
	greetPrinter(convertToUppercase, "Hi?")

	greetCreator := createGreetInTurkish
	greetPrinter(greetCreator, "Lara")

	func(name string) {
		greeting := "Selam " + name + " :)"
		fmt.Printf("%s\n", greeting)
	}("Mary")

	closure := func(name string) {
		greeting := "Selam " + name + " :)"
		fmt.Printf("%s\n", greeting)
	}
	closure("Ali")
	anotherGreetPrinter(closure, "Veli")
}

func createGreetInTurkish(name string) string {
	return "Selam " + name + " :)"
}

func createGreetInEnglish(name string) string {
	return "Hi " + name + " :)"
}

func convertToUppercase(arg string) string {
	return strings.ToUpper(arg)
}

func greetPrinter(function func(it string) string, name string) {
	var greeting = function(name)
	fmt.Printf("%s\n", greeting)
}

func anotherGreetPrinter(function func(it string), name string) {
	function(name)
}

/*func add(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}*/
