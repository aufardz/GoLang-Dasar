package main

import "fmt"

func main() {
	//deklarasi panjangnya
	// var person map[string]string = map[string]string
	//deklarasi singkat
	person := map[string]string{
		"name":    "Aufar",
		"address": "Bandung",
	}
	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Aufar Dzulfiqar"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
