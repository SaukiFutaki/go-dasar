package main

import "fmt"

func main() {
	//[index] lalu diikuti dengan tipe data yang akan diisi
	var arrayName [5]string

	arrayName[0] = "Sauki"
	arrayName[1] = "Futaki"
	arrayName[2] = "Ecca"
	arrayName[3] = "Laila"
	arrayName[4] = "Izza"

	fmt.Println(arrayName)

	//jika sudah diisi, maka tidak bisa diisi lagi
	//jika ada yg tidak diisi, maka akan diisi dengan nilai default yaitu 0

	//jika ingin menggunakan array tanpa menentukan jumlah index, maka bisa menggunakan [...]tipeData
	var angka = [...]int{
		90,
		95,
		100,
		85,
		80,
	}
	fmt.Println(angka)
	fmt.Println(len(angka))

	//mengganti nilai array index 0 yang awalnya 90 menjadi 100
	angka[0] = 100
	fmt.Println(angka)
}
