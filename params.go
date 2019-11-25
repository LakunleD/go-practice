package main

import "fmt"

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	values := []int{1, 2, 3, 4} //pass by reference
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10 //pass by value
	double(val)
	fmt.Println(val)

	doublePtr(&val) //pass by reference
	fmt.Println(val)
}
