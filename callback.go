package main

import "fmt"

func CallbackEvenNumberExample(limit int, f func(int)) {
	for i := 0; i < limit; i++ {
		if i%2 == 0 {
			f(i)
		}
	}
}

func main() {
	CallbackEvenNumberExample(10, func(i int) {
		fmt.Println(i)
	})
}
