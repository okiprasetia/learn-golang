package main

import "fmt" 

func main (){
	buah := [3]string{"Apel", "Pisang", "Jeruk"}
  // Mencetak elemen pertama (indeks 0)
  fmt.Println("Buah pertama:", buah[0]) 

  // Mencetak elemen kedua (indeks 1)
  fmt.Println("Buah kedua:", buah[1])

  // Mengubah nilai elemen ketiga (indeks 2)
  buah[2] = "Mangga"
  fmt.Println("Buah ketiga:", buah[2])
}