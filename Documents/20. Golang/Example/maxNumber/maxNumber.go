package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	num1 := 10
	num2 := 20

	result := max(num1, num2)
	fmt.Printf("Max Number %d and %d as: %d\n", num1, num2, result)

}
