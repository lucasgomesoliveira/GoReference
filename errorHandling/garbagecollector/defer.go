package main

import "fmt"

func cleanUp(name string) {

	fmt.Printf("Cleaning up %s\n", name)

}

func worker() {
	defer cleanUp("A")
	defer cleanUp("B")
	fmt.Println("Worker")
}

func main() {

	worker()

}
