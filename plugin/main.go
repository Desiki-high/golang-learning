package main

import (
	"fmt"
	"plugin"
)

type CalculatorPlugin interface {
	Operate(a, b int) int
}

func main() {
	var operator string
	fmt.Print("Enter operator (add or multiply): ")
	fmt.Scan(&operator)

	p, err := plugin.Open("./" + operator + "/plugin.so")
	if err != nil {
		fmt.Println("Error opening plugin:", err)
		return
	}

	symbol, err := p.Lookup("Calculator")
	if err != nil {
		fmt.Println("Error looking up symbol:", err)
		return
	}

	calcPlugin, ok := symbol.(CalculatorPlugin)
	if !ok {
		fmt.Println("Invalid plugin interface")
		return
	}

	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	result := calcPlugin.Operate(a, b)
	fmt.Printf("Result: %d\n", result)
}
