package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	addres := Address{"Semarang", "Jawa Tengah", ""}
	ChangeCountryToIndonesia(&addres)

	fmt.Println(addres)

}
