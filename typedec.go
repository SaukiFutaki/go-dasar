package main

import "fmt"

func typedec() {
	type tipeNama string

	var nama tipeNama = "Sauki"

	var namaTeman string = "Futaki"
	var namaTeman2 tipeNama = tipeNama(namaTeman)
	fmt.Println(nama)
	fmt.Println(namaTeman2)
}
