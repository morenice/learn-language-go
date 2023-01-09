package main

import (
	"fmt"
	"reflect"
)

func SliceArrayExample() {
	a1 := [10]string{"1", "2", "3"}
	s1 := []string{"1", "2", "3"}

	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(reflect.TypeOf(s1))

	fmt.Printf("len:%d cap:%d string(a1):%v\n", len(a1), cap(a1), a1)
	fmt.Printf("len:%d cap:%d string(s1):%v\n", len(s1), cap(s1), s1)

	//
	//append(a1, "4")
	s2 := append(s1, "4", "5")
	fmt.Printf("len:%d cap:%d string(s1):%v\n", len(s1), cap(s1), s1)
	fmt.Printf("len:%d cap:%d string(s2):%v\n", len(s2), cap(s2), s2)

	s1[0] = "a"
	fmt.Printf("len:%d cap:%d string(s1):%v\n", len(s1), cap(s1), s1)
	fmt.Printf("len:%d cap:%d string(s2):%v\n", len(s2), cap(s2), s2)

	s3 := s2[:3]
	fmt.Printf("len:%d cap:%d string(s2):%v\n", len(s2), cap(s2), s2)
	fmt.Printf("len:%d cap:%d string(s3):%v\n", len(s3), cap(s3), s3)

	s2[0] = "*"
	fmt.Printf("len:%d cap:%d string(s3):%v\n", len(s3), cap(s3), s3)
	fmt.Printf("len:%d cap:%d string(s1):%v\n", len(s1), cap(s1), s1)
	fmt.Printf("len:%d cap:%d string(s2):%v\n", len(s2), cap(s2), s2)

	fmt.Printf(s1[100])
}

func SliceMapExample() {

	var mymap map[string]string = make(map[string]string)
	mymap["1"] = "asdf"
	fmt.Println(mymap)

	var mymap2 = map[string]string{"123": "a"}
	fmt.Println(mymap2)

	mymap2["222"] = "b"
	mymap2["333"] = "c"

	elem := mymap2["222"]
	fmt.Println(elem)

	// key/value 조회 확인
	elem, exists := mymap2["222"]
	fmt.Println(exists)

	delete(mymap2, "222")

	delete(mymap2, "222")
	delete(mymap2, "222")
	delete(mymap2, "222")
	delete(mymap2, "222")
	elem2 := mymap2["222"]
	elem3 := mymap2["5555"]
	fmt.Println(elem2)
	fmt.Println(elem3)
}
