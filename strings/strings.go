package main

import "fmt"

func main() {

	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// uint8 is a byte

	// slicing strings
	fmt.Println(book[4:11])
	fmt.Println(book[4:])
	fmt.Println(book[:11])

	// multi line
	poem := `
	         roses are red the sky is blue
			 nobody will ever 
			 posibble love you`

	fmt.Println(poem)

}
