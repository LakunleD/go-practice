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

	y := 2

	switch y {
	case 1:
		fmt.Println("y is 1")
	case 2:
		fmt.Println("y is 2")
	case 3:
		fmt.Println("y is 3")
	default:
		fmt.Println("y is unknown")
	}

	z := 5
	switch {
	case z < 1:
		fmt.Println("z is less than 1")
	case z > 10:
		fmt.Println("z is greater than 10")
	default:
		fmt.Println("z is between 1 and 10")
	}
}
