package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
)

func main()  {
	type Person struct {
		Name string
		Age  int
	}

	data := map[string]interface{}{
		"name": "Bob",
		"age":  20,
	}

	var p Person
	err := mapstructure.Decode(data, &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s is %d years old.\n", p.Name, p.Age)  // Output: Bob is 20 years old.
}
