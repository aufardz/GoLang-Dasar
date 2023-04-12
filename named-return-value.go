package main

import "fmt"

func getFullName() (firstname, middlename, lastname string) {
	firstname = "Aufar"
	middlename = "Bin"
	lastname = "Dzulfiqar"
	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a, b, c)

}
