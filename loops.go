package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("+++++++++++++++++++++")

	for i := 0; i < 5; i++ {
		if i > 2 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("+++++++++++++++++++++")

	for i := 0; i < 5; i++ {
		if i < 2 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("+++++++++++++++++++++")

	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}
}
