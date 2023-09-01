package main

type AddPlugin struct{}

func (a AddPlugin) Operate(x, y int) int {
	return x + y
}

var Calculator AddPlugin
