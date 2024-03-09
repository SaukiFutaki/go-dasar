package main

import "fmt"

func aww() interface{} {
	return "awwwww"
}

func wew() any {
	return 1

}

func main() {
	wow := wew()
	awwwww := aww()
	fmt.Println(awwwww, wow)
}
