package main

import "fmt"

func main() {
	people := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}

	fmt.Println("Info User")
	for name, age := range people {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

}
