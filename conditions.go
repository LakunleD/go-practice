package main

import "fmt"

func main() {
	x := 20
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	if x > 100 {
		fmt.Println("x is greater than 100")
	} else {
		fmt.Println("x is less than 100")
	}
}
