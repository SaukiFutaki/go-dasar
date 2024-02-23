package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println("Aplikasi selesai")
	fmt.Println("Pesan error:", message)
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}

}
func main() {
	runApp(true)

}
