package main

import "fmt"

// Fungsi untuk menghitung penjumlahan dua angka
func jumlahkan(angka1 int, angka2 int) int {
  hasil := angka1 + angka2
  return hasil
}

func main() {
  angka1 := 5
  angka2 := 3

  // Memanggil fungsi jumlahkan dan menyimpan hasilnya
  total := jumlahkan(angka1, angka2)

  fmt.Println("Hasil penjumlahan:", total) // Output: Hasil penjumlahan: 8
}