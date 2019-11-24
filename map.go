package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"AMZN": 122.9,
		"GOOG": 178.19,
		"MSFT": 98.12,
	}
	fmt.Println(len(stocks))

	fmt.Println("+++++++++++++++++++++")
	fmt.Println(stocks["MSFT"])
	fmt.Println("+++++++++++++++++++++")
	fmt.Println(stocks["TSLA"])
	fmt.Println("+++++++++++++++++++++")

	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}
	fmt.Println("+++++++++++++++++++++")

	stocks["TSLA"] = 34.2
	fmt.Println(stocks)
	fmt.Println("+++++++++++++++++++++")

	delete(stocks, "AMZN")
	fmt.Println(stocks)
	fmt.Println("+++++++++++++++++++++")

	for key := range stocks {
		fmt.Println(key)
	}
	fmt.Println("+++++++++++++++++++++")

	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}

}
