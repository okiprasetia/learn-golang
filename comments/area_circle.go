package main

import (
	"fmt"
	"math"
)

func main()  {
	// Declare a variable to store the radius of the circle
	var radius float64
	// Prompt the user to enter the radius of the circle
	fmt.Print("Enter the radius of circle : ")
	// Read the radius entered by the user from the console
	fmt.Scanln(&radius)
	// calculate the area of the circle using the formula : Area = π * r^2
	area := math.Pi * math.Pow(radius, 2)
	// Print the calculated area of the circle
	fmt.Printf("the area of the circle with radius %.2f is %.2f\n", radius, area)

}