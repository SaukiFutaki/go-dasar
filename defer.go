package main

import "fmt"

func loggin() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer loggin()
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
