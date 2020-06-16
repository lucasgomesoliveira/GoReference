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

	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)

	// functions without pointers doesn't change in this scope
	val := 10
	double(val)
	fmt.Println("func accesing directly -> ", val)

	// functions with pointers are changed in the scope of declaration
	doublePtr(&val)
	fmt.Println("func accessing by pointers -> ", val)

}
