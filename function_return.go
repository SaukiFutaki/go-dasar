package main

import "fmt"

func getHello(name string, age int) string {
	return "Hello " + name + age

}
func main() {
	result := getHello("John", 20)
	fmt.Println(result)
}
