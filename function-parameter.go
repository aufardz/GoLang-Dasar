package main

import "fmt"

func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)

}

func main() {
	firstname := "Aufar"
	sayHelloTo(firstname, "Dzulfiqar")
	sayHelloTo("Anneke", "Putri")
}
