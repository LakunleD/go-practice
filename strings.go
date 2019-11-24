package main

import "fmt"

func main() {
	book := "Half of a yellow sun"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	//slice (start, end)
	fmt.Println(book[10:16])
	//slice(no end)
	fmt.Println(book[8:])
	//slice(no start)
	fmt.Println(book[:8])
}
