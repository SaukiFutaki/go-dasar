package main

import "fmt"

// nill hanya bisa digunakan pada tipe data pointer, interface, map, slice, channel, dan fungsi
func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	test1 := newMap("soqi")
	test2 := newMap("")
	fmt.Println(test1)
	fmt.Println(test2)

	if test2 == nil {
		fmt.Println("test2 is nil")
	} else {
		fmt.Println("test2 is not nil")
	}
}
