package main

import "fmt"

//parameter
func sayHello(firstName string, lastName string, age int) {
	fmt.Println("Halo nama saya", firstName, lastName, "umur saya", age, "tahun")
}

//return value
//func namaFunction(parameter) tipeDataReturn
func gethello(p string) string {
	return "Hallo " + p
}

//return value multiple
//func namaFunction(parameter) (tipeDataReturn1, tipeDataReturn2)
func getFullName() (string, string) {
	return "John", "Doe"
}

func main() {

	sayHello("John", "Doe", 20)

	fmt.Println(gethello("John"))

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
