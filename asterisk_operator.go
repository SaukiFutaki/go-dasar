package main

import "fmt"

type Addres struct {
	City, Province, Country string
}

func main() {
	var addres1 = Addres{"Semarang", "Jawa Tengah", "Indonesia"}
	var addres2 *Addres = &addres1 //pointer

	fmt.Println(addres1) //semua data
	fmt.Println(addres2) //alamat memori

	addres2.City = "Jakarta"
	fmt.Println(addres1) //ikut berubah
	fmt.Println(addres2) //berubah menjadi jakarta

	addres2 = &Addres{"Malang", "Jawa Timur", "Indonesia"} //pointer baru
	fmt.Println(addres1)                                   //tidak ikut berubah
	fmt.Println(addres2)                                   //berubah menjadi malang

	//jika kita menggunakan * diawal variabel pointer maka kita akan mendapatkan data asli
	// *addres2 = Addres{"Surabaya", "Jawa Timur", "Indonesia"} //mengubah data asli yaitu addres1 menjadi addres2
	// fmt.Println(addres2)
	// fmt.Println(addres1) //ikut berubah
}
