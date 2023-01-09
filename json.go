package main

import (
	"encoding/json"
	"fmt"
)

// public member variable case
//   - first char capital case
type Address struct {
	Id          int    `json:"id"`
	FullName    string `json:"fullName"`
	Name0       string `json:"name0"`
	Code0       string `json:"code0"`
	Description string
}

// private member variable case
//   - first char lower case
type Address2 struct {
	id       int    `json:"id"`
	fullName string `json:"fullName"`
}

func unmarshal() {
	inputData := `{
	"id": 100,
	"fullName": "address full name",
	"name0": "first",
	"code0": "A1823"
 }`

	out := Address{}
	json.Unmarshal([]byte(inputData), &out)
	fmt.Printf("%#v\n", out)

	// Don't use private member variable
	// During unmarshal Address2 structure don't fill.
	out2 := Address2{}
	json.Unmarshal([]byte(inputData), &out2)
	fmt.Printf("%#v\n", out2)
}

func main() {
	unmarshal()
}
