package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool
	var noKtpAufar NoKtp = "3273132710000006"
	fmt.Println(noKtpAufar)
	var status Married = true
	fmt.Println(status)
}
