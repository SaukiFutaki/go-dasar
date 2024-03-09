package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	joko := Customer{
		Name:    "Joko",
		Age:     22,
		Address: "Jakarta",
	}

	fmt.Println(joko)

	Budi := Customer{"budi", "bandung", 22}

	fmt.Println(Budi)
}
