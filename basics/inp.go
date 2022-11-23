package main

import "fmt"

func main() {
	fmt.Println("hello")

	var name string
	var age int
	var isPass bool

	fmt.Println("Enter name, age, isPass")
	// fmt.Scanf("%s %d %t", &name, &age, &isPass)
	fmt.Println(name, age, isPass)

	fmt.Println("Enter no of goals scored")
	var goals int
	// fmt.Scan(&goals)
	fmt.Println("goals scored: ", goals)
}
