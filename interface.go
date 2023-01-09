package main

import "fmt"

type myinterface interface {
	myfunction() int
}

type MyInt int

func (rcv MyInt) myfunction() int {
	return 0
}

type MyString string

func (rcv MyString) myfunction() int {
	return 1
}

func TestingInterface() {
	var a1 = []myinterface{MyInt(3), MyString("aaa")}
	for _, i := range a1 {
		fmt.Println(i.myfunction())
	}
}
