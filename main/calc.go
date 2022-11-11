package main

import "fmt"

/* Adds two numbers */
func add(x, y int) int {
	return x + y
}

// Subtracts two numbers
func sub(x, y int) int {
	return x - y
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(add(15, 30))
}
