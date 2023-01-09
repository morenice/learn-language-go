package main

import (
	"fmt"
	"strings"
)

// return variable case #1
func GetFirstAndLastChar(s string) (string, string) {
	splitedString := strings.Split(s, "")
	return splitedString[0], splitedString[len(splitedString)-1]
}

// return variable case #2
// named return variable
func GetFirstAndLastChar2(s string) (first string, last string) {
	splitedString := strings.Split(s, "")
	first = splitedString[0]
	last = splitedString[len(splitedString)-1]

	// only write return keyword
	return
}

// return variable case #3 with defer
// useful file/socket close that defer
func GetFirstAndLastChar3(s string) (first string, last string) {
	defer fmt.Println("Run after to finish this function")
	defer fmt.Println("Run2 after to finish this function")

	splitedString := strings.Split(s, "")
	first = splitedString[0]
	last = splitedString[len(splitedString)-1]
	return
}

// dynamic argumens
// s: string type list
func PrintAllParams(s ...string) {
	fmt.Println(s)
}

func FunctionValueExample() {
	func1 := func(x int, y int) int {
		return x + y
	}
	fmt.Println(func1(1, 2))
}

func main() {
	FunctionValueExample()
}
