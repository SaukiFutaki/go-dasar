package main

func conversion() {
	var nilai32 int32 = 32768
	//konversi nilai32 ke int64
	var nilai64 int64 = int64(nilai32)
	//konversi nilai32 ke int16
	var nilai16 int16 = int16(nilai32)

	println(nilai32)
	println(nilai64)
	println(nilai16)

	var name = "Sauki"
	println(name[0])
	var e uint8 = name[0]
	var eString = string(e)

	println(eString)
}
