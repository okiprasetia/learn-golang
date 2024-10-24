package main

import (
	"fmt"

	"maths/mathutils"
)

// mendapatkan hasil perkalian dari fungsi yang diimport dari package mathutils
func main() {
	hasil := mathutils.Multiply(2, 3)
	fmt.Println("hasil dari perkalian : ", hasil)
}
