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
	var angka = [3]int{
		90,
		95,
		100,
	}
	fmt.Println(angka)
}
