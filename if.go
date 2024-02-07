package main

import "fmt"

func main() {
	name := "Doe"

	if name == "John" {
		fmt.Println("Hello John")
	} else if name == "Doe" {
		fmt.Println("Hello Doe")
	} else {
		fmt.Println("Hello Stranger")
	}

	//short statement
	if panjang := len(name); panjang > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
