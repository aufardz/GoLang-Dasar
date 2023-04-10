package main

import "fmt"

func main() {
	var name string

	name = "Aufar Dzulfiqar"
	fmt.Println(name)
	// name = "Dzulfiqar"
	// fmt.Println(name)
	var age = 23
	fmt.Println("Umur = ", (age))
	birthday := "Bandung"
	fmt.Println("Kota Lahir = ", (birthday))
	fmt.Println("")

	var gfname = "Anneke Putri"
	fmt.Println(gfname)
	// gfname = "Putri"
	// fmt.Println(gfname)
	var agegf = 22
	fmt.Println("Umur = ", (agegf))
	birthday = "Padang"
	fmt.Println("Kota Lahir = ", (birthday))

	var (
		firstname = "Jonas"
		lastname  = "Hofmann"
	)
	fmt.Println("")
	fmt.Println(firstname)
	fmt.Println(lastname)
}
