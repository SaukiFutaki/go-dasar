package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		//jika i == 5, kita stop perulangan
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}
}
