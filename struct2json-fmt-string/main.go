package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Test struct {
	Test  string      `json:"test"`
	Ts    int         `json:"ts"`
	Child []TestChild `json:"child"`
}

type TestChild struct {
	ChildString string `json:"childString"`
	ChildTime   int    `json:"childTime"`
}

func main() {
	var testChildren []TestChild
	testChildren = append(testChildren, TestChild{ChildString: "this is the first child", ChildTime: 01234})
	testChildren = append(testChildren, TestChild{ChildString: "this is the second child", ChildTime: 56789})

	test := Test{
		Test:  "this is a test string",
		Ts:    1677045875836,
		Child: testChildren,
	}

	println(test.String())
}

func (args *Test) String() string {
	b, err := json.Marshal(*args)
	if err != nil {
		return fmt.Sprintf("%+v", *args)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *args)
	}
	return out.String()
}
