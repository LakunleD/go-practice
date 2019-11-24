package main

import "fmt"

func main() {
	numbers := []int{12, 4, 5, 34, 6, 7, 8}
	max := numbers[0]

	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}
	fmt.Println(max)
}
