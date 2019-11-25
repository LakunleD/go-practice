package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func main() {
	t1 := Trade{"MSFT", 10, 99.91, true}

	fmt.Println(t1)

	fmt.Printf("%+v\n", t1)

	fmt.Println(t1.Symbol)

	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.91,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)
	t3.Symbol = "MSFT"
	t3.Volume = 10
	t3.Price = 99.91
	t3.Buy = true
	fmt.Printf("%+v\n", t3)
}
