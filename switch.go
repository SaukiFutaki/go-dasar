package main

func main() {
	name := "Doe"

	switch name {
	case "John":
		println("Hello John")
	case "Doe":
		println("Hello Doe")
	default:
		println("Hello Stranger")
	}

	//short statement
	switch panjang := len(name); panjang > 5 {
	case true:
		println("Terlalu panjang")
	case false:
		println("Nama sudah benar")
	}

	//switch tanpa kondisi
	pendek := len(name)
	switch {
	case pendek > 5:
		println("Terlalu panjang")
	case pendek < 5:
		println("Nama Pendek")
	default:
		println("default")
	}
}
