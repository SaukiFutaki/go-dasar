package main

import "fmt"

// function as parameter
// function bisa dijadikan parameter
// function sayHelloWithFilter memiliki 2 parameter, yaitu name dan filter
// filter merupakan function yang memiliki 1 parameter string dan mengembalikan string
func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

// function spamFilter memiliki 1 parameter string dan mengembalikan string
func spamFilter(name string) string {
	if name == "ANJING" {
		return "..."
	} else {
		return name
	}
}

func main() {
	filter := spamFilter
	sayHelloWithFilter("ANJING", filter)
}
