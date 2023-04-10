package main

import "fmt"

func main() {
	//Array Manual
	var names [3]string

	names[0] = "Aufar"
	names[1] = "Dzulfiqar"
	names[2] = "Anput"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Array Langsung
	var values = [3]int{
		27,
		10,
		00,
	}
	fmt.Println(values)
	fmt.Println(len(names))
	fmt.Println(len(values))
}
