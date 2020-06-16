package main

import (
	"fmt"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func main() {

	t1 := Trade{"MSFT", 10, 99.98, true}

	fmt.Println(t1)
	fmt.Printf("%+v\n", t1)
	fmt.Println(t1.Symbol)

	t2 := Trade{Symbol: "MSFT", Volume: 10, Price: 99.98, Buy: true}

	fmt.Println(t2)
	fmt.Printf("%+v\n", t2)
	fmt.Println(t2.Symbol)

	t3 := Trade{}
	fmt.Println(t3)

}
