package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		//jika i genap, kita skip
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}
