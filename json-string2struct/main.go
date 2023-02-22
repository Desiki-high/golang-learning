package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	Str := "{\"name\":\"xxx\",\"age\":12}"
	fmt.Printf("%+v\n", GetStruct(Str))
}

func GetStruct(s string) Person {
	var p Person
	err := json.Unmarshal([]byte(s), &p)
	if err != nil {
		panic(err)
	}
	return p
}
