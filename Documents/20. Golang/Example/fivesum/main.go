package main

import "fmt"

func main() {
	fmt.Println("Bang cuu chuong 5:")
	for i := 1; i <= 10; i++ {
		result := 5 * i
		fmt.Printf("5 x %d = %d\n", i, result)
	}
}
