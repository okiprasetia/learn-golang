package main

import "fmt"
import "math"

/*
func calculate(d float64)(float64, float64){
	//hitung luas
	area := math.Pi * math.Pow(d / 2, 2)

	//hitung keliling
	circumfrence := math.Pi * d

	//kembalikan 2 nilai
	return area, circumfrence
}
*/
func calculate(d float64)(area float64, circumfrence float64) {
	area = math.Pi * math.Pow(d / 2, 2)
	circumfrence = math.Pi * d
	
	return
}

func main(){
	var diameter float64 = 15
	var area, circumfrence = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %2f \n", area)
	fmt.Printf("keliling lingkaran\t: %2.f \n", circumfrence)
}