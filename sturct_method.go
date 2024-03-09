package main

import "fmt"

type Customer struct {
	Name, Address string
}

type Point struct {
	X, Y int
}

// method adalah function yang menempel pada struct
// method adalah function

func (p Point) multiply() float32 {
	return float32(p.X * p.Y)
}

func (customer Customer) sayHelloCustomer() {
	fmt.Println("Hello", customer.Name, "from", customer.Address)
}

func main() {
	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
	}

	testOne := Point{3, 4}
	fmt.Println(testOne.multiply())
	joko.sayHelloCustomer()
}
