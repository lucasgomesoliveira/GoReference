package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)

	}

	for i := 0; i < 3; i++ {

		if i > 1 {
			continue
			fmt.Println("pass")
		}
		fmt.Println(i)

	}

}
