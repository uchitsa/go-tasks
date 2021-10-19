package main

import "fmt"

func main() {
	var stack []string

	stack = append(stack, "Two")
	stack = append(stack, "One")

	for len(stack) > 0 {
		n := len(stack) - 1
		fmt.Println(stack[n])
		stack = stack[:n]
	}
}
