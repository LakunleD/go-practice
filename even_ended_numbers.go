/*
An "even ended number" is a number whose first and last digit are the same.
count the number of even ended numbers that are a multiplication of 4 digit numebrs.
*/
package main

import "fmt"

func main() {
	count := 0

	for i := 1000; i <= 9999; i++ {
		for j := i; j < 9999; j++ {
			number := i * j

			number_string := fmt.Sprintf("%d", number)

			if number_string[0] == number_string[len(number_string)-1] {
				count++
			}
		}
	}
	fmt.Println(count)
}
