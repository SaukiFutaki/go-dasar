package main

import "fmt"

func sayGoodBye(name string) string {
	return "Good Bye " + name

}

func main() {
	//jika kita ingin menyimpan fungsi sayGoodBye ke dalam variabel jangan panggil menggunakan
	//kurung buka dan kurung tutup seperti ini sayGoodBye()
	//karena ini akan memanggil fungsi sayGoodBye dan bukan menyimpan fungsi sayGoodBye ke dalam variabel
	goodBye := sayGoodBye
	fmt.Println(goodBye("John"))
}
