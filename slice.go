package main

import "fmt"

func main() {
	iniArray := [...]string{"nol", "satu", "dua", "tiga", "empat", "lima"}

	//1 adalah index awal, 4 adalah index akhir
	slice1 := iniArray[1:4]

	//iniArray[:3] artinya mengambil elemen dari index 0 sampai 2
	slice2 := iniArray[:3]

	//iniArray[3:] artinya mengambil elemen dari index 3 sampai terakhir
	slice3 := iniArray[3:]

	//iniArray[:] artinya mengambil semua elemen
	slice4 := iniArray[:]

	fmt.Println(iniArray)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(cap(slice1))

}
