package main

func main() {
	var nilai = 80
	var absensi = 80

	var lulusNilai = nilai >= 75
	var lulusAbsensi = absensi >= 75

	var lulus = lulusNilai && lulusAbsensi
	println(lulus)
}
