package main

import "fmt"

// interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
// sebuah interface berisikan definisi-definisi method
// biasanya interface digunakan sebagai kontrak

type Name interface {
	//GetName adalah method yang harus diimplementasikan oleh sebuah struct agar disebut sebagai implementasi dari interface Name
	GetName() string
}

func sayHello(name Name) {
	fmt.Println("Hello", name.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	soqi := Person{"soqi"}
	sayHello(soqi)
}
