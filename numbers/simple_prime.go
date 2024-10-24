package main

import "fmt"

// fungsi mencari bilangan prima dengan sisa hasil bagi
// bilangan prima akan selalu sisa 0 sedangkan bilangan ganjil tidak
func IsPrime(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}

// masukan sebuah angka lalu panggil fungsi untuk IsPrime
// cetak hasil bilangan tersebut bilangan prima atau bukan
func main() {
	num := 10
	hasil := IsPrime(num)
	if hasil == true {
		fmt.Printf("bilangan %d adalah bilangan prima", num)
	} else {
		fmt.Printf("bilangan %d bukan bilangan prima", num)
	}
}
