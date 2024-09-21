package main

import "fmt"

func menyambung_huruf(nama_depan string, nama_belakang string) string {
	nama_lengkap := nama_depan +" "+nama_belakang
	return nama_lengkap
}

func main (){
	nama_depan := "james"
	nama_belakang := "bond"

	nama := menyambung_huruf(nama_depan, nama_belakang)
	fmt.Println("nama lengkap anda :", nama)
}