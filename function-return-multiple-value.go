package main

import "fmt"

func getFullName() (string, string) {
	return "Aufar", "Dzulfiqar"
}

func main() {
	firstname, lastname := getFullName()
	fmt.Println(firstname, lastname)
}
