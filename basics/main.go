package main

import "fmt"

func main() {
	var name string = "Sabuj"
	const ttl = 1000
	fmt.Println(name, ttl)

	fmt.Print(name)
	fmt.Print(ttl)
	fmt.Println(name, ttl)
	fmt.Printf("Age: %d", 20)

	var numr float64 = 65000.5
	var age int = 23
	fmt.Printf("\nAnnual salary: %g", numr)
	fmt.Printf("\nAnnual salary: %f", numr)
	fmt.Printf("\n%s is %d years old", name, age)

}
