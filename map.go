package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "John",
		"address": "Los Angeles",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//make untuk membuat map baru dengan tipe data key string dan value string
	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang" //menambahkan key title dengan value Belajar Go-Lang
	book["author"] = "John Doe"       //menambahkan key author dengan value John Doe
	book["ups"] = "Salah"             //menambahkan key ups dengan value Salah

	fmt.Println(book) //map[title:Belajar Go-Lang author:John Doe ups:Salah]

	delete(book, "ups") //menghapus key ups

	fmt.Println(book)      //map[title:Belajar Go-Lang author:John Doe]
	fmt.Println(len(book)) //2
}
