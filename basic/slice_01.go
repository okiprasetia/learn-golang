package main

import "fmt"

func printArray (list []string){
	for i := 0; i < len(list); i++ {
		fmt.Printf("elemen %d : %s\n", i, list[i])
	}
}

func main (){
	project_name := []string{"project-01", "project-02", "project-03"}
	printArray(project_name)
}