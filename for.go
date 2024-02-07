package main

import "fmt"

func main() {
	counter := 0

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	fmt.Println("Selesai perulangan")

	//for dengan statement
	for inc := 0; inc <= 10; inc++ {
		fmt.Println("Perulangan ke ", inc)

	}
	fmt.Println("Selesai perulangan")

	//for range
	names := []string{"John", "Wick", "Ethan", "Hunt"}
	//jika kita tidak membutuhkan index, kita bisa menggunakan _ (underscore)
	// for _, name := range names {
	for i, name := range names {
		fmt.Println("Index ke ", i, "=", name)
	}

}
