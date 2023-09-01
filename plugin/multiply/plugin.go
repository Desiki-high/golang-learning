package main

type MultiplyPlugin struct{}

func (m MultiplyPlugin) Operate(x, y int) int {
	return x * y
}

var Calculator MultiplyPlugin
