package main

import "fmt"

type Addres struct {
	City, Province, Country string
}

func main() {
	// pointer adalah reference ke lokasi data
	// pointer adalah tipe data
	// pointer adalah alamat memori
	// pointer adalah data yang menunjuk data lain

	//contoh dibawah adalah cara pemborosan memori
	//karena data yang di copy adalah data yang sama
	//bukan adddres dua adalah addres satu
	//tapi addres2 menduplikasi addres1
	addres1 := Addres{"Semarang", "Jawa Tengah", "Indonesia"}
	addres2 := addres1

	addres2.City = "Jakarta"

	fmt.Println(addres1)
	fmt.Println(addres2)

	//cara mengatasi pemborosan memori dengan menggunakan pointer
	addres3 := Addres{"Semarang", "Jawa Tengah", "Indonesia"}

	//& adalah operator untuk mendapatkan addres
	addres4 := &addres3

	addres4.City = "Jakarta"

	fmt.Println(addres3)
	fmt.Println(addres4)

}
