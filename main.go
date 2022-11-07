package main

import (
	"fmt"

	"github.com/JanaSabuj/playground/sub"
	"github.com/go-modules-by-example/submodules/b"
	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("Hi I am Sabuj")

	// marshal
	marshal, err := yaml.Marshal(
		map[string]string{
			"Name":    "Sabuj",
			"Company": "Flipkart",
			"Edu":     "JU",
		},
	)
	fmt.Println(string(marshal), err)
	fmt.Println(b.Name)

	// submodule sub
	fmt.Println(sub.GreetMsg)
}
