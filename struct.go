package main

import "fmt"

type Bio struct {
	// struct adalah kumpulan dari field atau property
	//best practice nama field diawali huruf besar
	//struct tidak bisa langsung digunkan, kita harus membuat instance dari struct
	//struct mirip destructor di javascript

	Name, Addres string
	Age          int
}

func main() {
	var soqi Bio
	soqi.Name = "Soqi"
	soqi.Addres = "Jakarta"
	soqi.Age = 22

	fmt.Println(soqi)
}
