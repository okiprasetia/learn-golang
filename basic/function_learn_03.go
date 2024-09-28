package main

import "fmt"
import "strings"

func printMessage(message string, arr []string){ //parameter message & slice
	var nameString = strings.Join(arr," ")
	fmt.Println(message, nameString)
}

func main (){
	names:= []string{"oki","prasetia"}
	printMessage("hallo", names) //argumen message (hallo) & slice string names 
}