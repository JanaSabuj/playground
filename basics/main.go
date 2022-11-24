package main

import "fmt"

func main() {
	var name = "Sabuj"
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

	// ops
	fmt.Println("\n", 10/3)
	fmt.Println(10.0 / 3.0)

	// Go doesn't support implicit type casting - it only does explicit type casting
	var f float32 = 4.5
	var ff = 5.5
	var no = int(f)
	var nono = int(ff)
	fmt.Println(no, nono)

}
