package main

import "fmt"

func main() {
	animals := []string{"monkey", "giraffe", "cat"}

	fmt.Printf("animals = %v (type %T)\n", animals, animals)

	fmt.Println(len(animals))

	fmt.Println("+++++++++++++++++++++")

	fmt.Println(animals[1])

	fmt.Println("+++++++++++++++++++++")

	fmt.Println(animals[1:])

	fmt.Println("+++++++++++++++++++++")

	for i, name := range animals {
		fmt.Printf("The name of the animal at position %d is %s\n", i, name)
	}

	fmt.Println("+++++++++++++++++++++")

	animals = append(animals, "dog")

	fmt.Println(animals)
}
