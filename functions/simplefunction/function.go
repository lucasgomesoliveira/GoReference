package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {

	fmt.Println(add(1, 2))

	div, mod := divmod(10, 7)

	fmt.Println(div, mod)

}
