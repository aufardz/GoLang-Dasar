package main

import "fmt"

func main() {
	var name = "Aufar"
	var name2 = "Anne"
	var result bool = name == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
