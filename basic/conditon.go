package main

import "fmt"

func main(){
	a := 0
	
	if a <= 80 && a >=70   {
		fmt.Println("nilai anda B")
	}else if a <= 70 && a >=30 {
		fmt.Println("nilai anda C")
	}else if a >=80 {
		fmt.Println("nilai anda C")
	}else {
		fmt.Println("tidak lulus")
	}
}