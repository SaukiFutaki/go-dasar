package main

func main() {
	var a = 100
	var b = 20
	var tambah = a + b
	var kurang = a - b
	var kali = a * b
	var bagi = a / b
	var modulo = a % b
	println(tambah)
	println(kurang)
	println(kali)
	println(bagi)
	println(modulo)

	//augmented assignment
	var i = 10

	//i = i + 10
	i += 20

	println(i)

	//unary operator
	var k = 10
	//k = k + 1
	k++
	println(k)
}
