package main

import (
	"fmt"
)

func EmptyArray() {
	var emptyArray1 []int
	var emptyArray2 []int = nil

	fmt.Println(emptyArray1)           // {}
	fmt.Println(emptyArray2)           // {}
	for _, item := range emptyArray2 { // skipped
		fmt.Println(item)
	}
}

func DeclareArray() {
	array1 := [10]string{"1", "2", "3"}
	fmt.Println(len(array1))

	array2 := [...]string{"1", "2", "3"}
	fmt.Println(len(array2))
}

func main() {
	EmptyArray()
	DeclareArray()
}
