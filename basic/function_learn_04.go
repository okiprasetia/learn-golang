package main

import "fmt"

func pembagian(a, b int){
	if b == 0 {
		fmt.Printf("invalid! %d tidak bisa dibagi oleh %d\n", a,b)
		return
	}

	hasil := a / b
	fmt.Printf("hasil bagi dari %d dibagi %d adalah %d\n", a, b, hasil)
}

func main(){
	pembagian(5,0)
	pembagian(10,5)
}