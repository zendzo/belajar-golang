package main

import "fmt"

func main(){
	var nilai = 80
	var absensi = 80

	var lulusNilai = nilai > 80
	var lulusAbsensi = absensi > 80

	var lulus =  lulusNilai && lulusAbsensi

	fmt.Println(lulus)
}
