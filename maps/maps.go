package main

import "fmt"

func main() {

	stocks := map[string]float64{
		"AMAZON": 1699.8,
		"GOOGLE": 1129.19,
		"MCSOFT": 98.61}

	fmt.Println(stocks)

	// get value
	fmt.Println(stocks["MCSOFT"])

	// if not found it brings 0
	fmt.Println(stocks["MCSOFT"])

	// test value
	value, ok := stocks["TESLA"]
	if !ok {
		fmt.Println("tesla not found")
	} else {
		fmt.Println(stocks)
	}

	fmt.Println(value)

	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}

}
