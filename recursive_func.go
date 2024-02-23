package main

import "fmt"

func recursiveFunc(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFunc(value-1)
	}

}

func main() {
	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(result)
	fmt.Println(recursiveFunc(10))
}
