package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan Ke : ", counter)
	// 	counter++
	// }

	// Init Statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke : ", counter)
	}

	slice := []string{"Aufar", "Dzulfiqar", "Anneke", "Putri", "Anput"}
	//Manual
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	// Simple
	// for _, value := range slice (Tanda Underscored hanya untuk menggunakan variable yang tidak dipakai)
	// fmt.Println(value) sehingga (outputnya seperti)

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}
	//untuk Map
	person := make(map[string]string)
	person["Name"] = "Aufar"
	person["Title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
