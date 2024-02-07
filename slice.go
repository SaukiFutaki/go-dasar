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
	fmt.Println("ini slice 1", slice1)
	fmt.Println("ini slice 2", slice2)
	fmt.Println("ini slice 3", slice3)
	fmt.Println("ini slice 4", slice4)

	//len() untuk mendapatkan panjang slice
	fmt.Println("panjang slice1 adalah", len(slice1))
	//cap() untuk mendapatkan kapasitas slice
	fmt.Println("kapasitas slice1 adalah", cap(slice1))

	//append() untuk menambahkan elemen pada slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	dayslice1 := days[5:]         //Sabtu dan Minggu
	dayslice1[0] = "Sabtu baru"   //mengubah elemen slice0 menjadi Sabtu Lagi
	dayslice1[1] = "Minggu baru"  //mengubah elemen slice1 menjadi Minggu Lagi
	fmt.Println("ini", dayslice1) //Sabtu Lagi dan Minggu Lagi

	dayslice2 := append(dayslice1, "Hari Baru") //menambahkan elemen pada slice
	fmt.Println("dayslice2 adalah", dayslice2)  //Sabtu Lagi, Minggu Lagi, Hari Baru
	fmt.Println("dayslice1 adalah", dayslice1)  //Sabtu Lagi, Minggu Lagi
	fmt.Println("ini array days", days)

	//make untuk membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "nol"
	newSlice[1] = "satu"
	//newSlice[2] = "dua" //error karena panjang slice hanya 2

	fmt.Println(newSlice)      //nol, satu
	fmt.Println(len(newSlice)) //2
	fmt.Println(cap(newSlice)) //5

	newSlice2 := append(newSlice, "dua")

	fmt.Println(newSlice2)      //nol, satu, dua
	fmt.Println(len(newSlice2)) //3
	fmt.Println(cap(newSlice2)) //5

	//copy untuk menyalin elemen dari slice ke slice lain
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice) //menyalin elemen dari fromSlice ke toSlice
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

}
