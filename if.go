package main

func main() {
	var name = "Aufar"
	if name == "Aufar" {
		println("Hallo Aufar")
	} else if name == "Anne" {
		println("Hallo Anne")
	} else {
		println("Hallo, Kamu siapa?")
	}

	// var length = len(name)
	//If Short Statements
	if length := len(name); length > 5 {
		println("Nama Terlalu Panjang")
	} else {
		println("Nama Sudah Benar")
	}
}
