package main

import "fmt"

func calculate(numbers ...int) float64{
	total := 0
	for _, number := range numbers {
		total += number
	}

	avg := float64(total) / float64(len(numbers))
	return avg
}

func main(){
	avg := calculate(2, 4, 3)
	msg := fmt.Sprintf("rata rata : %.2f", avg)
	fmt.Println(msg)
}