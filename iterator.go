package main

import (
	"fmt"
)

func ForRangeSample() {
	var sum int = 0
	var values = []int{1, 2, 3, 5}

	for _, v := range values {
		sum += v
	}
	fmt.Println("sum:", sum)
}

func ForLoopSample() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
