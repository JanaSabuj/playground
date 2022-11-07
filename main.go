package main

import (
	"fmt"

	"github.com/JanaSabuj/playground/morning"
	"github.com/JanaSabuj/playground/sub"           // subpackage
	"github.com/JanaSabuj/playground/subm"          // submodule
	"github.com/go-modules-by-example/submodules/b" // 3rd-party submodule
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

	// subpackage
	fmt.Println(sub.GreetMsg)
	fmt.Println(sub.AlohaMsg)

	// submodule
	fmt.Println(subm.GreetMsg)
	fmt.Println(subm.NewMsg)

	// submodule
	fmt.Println(morning.GreetMsg)
}
