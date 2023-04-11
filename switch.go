package main

import "fmt"

func main() {
	name := "Aufar"

	switch name {
	case "Aufar":
		fmt.Println("Halo Aufar")
	case "Anneke":
		fmt.Println("Halo Anneke")
	default:
		fmt.Println("Halo, Kamu siapa?")

	}

	// // switch short statement
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama Sudah Benar")
	// 	// default
	// 	// fmt.Println("Nama")
	// }

	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")

	}

}
