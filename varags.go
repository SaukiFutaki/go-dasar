package main

import "fmt"

//fungsi sumAll menerima parameter numbers yang merupakan varargs
//varargs adalah parameter yang bisa menerima lebih dari satu argumen
//numbers ...int artinya numbers adalah varargs yang menerima argumen bertipe int
func sumAll(numbers ...int) int {
	total := 0
	//menggunakan for range untuk menjumlahkan semua elemen numbers

	for _, number := range numbers {
		total += number
	}
	return total

}
func main() {
	total := sumAll(10, 10, 10, 10, 10, 10, 10, 10)
	fmt.Println(total)
	//jika kita ingin mengirimkan slice ke dalam fungsi sumAll
	slice := []int{10, 10}
	//maka slice harus dipecah menjadi elemen elemen dengan menggunakan ...
	fmt.Println(sumAll(slice...))

}
