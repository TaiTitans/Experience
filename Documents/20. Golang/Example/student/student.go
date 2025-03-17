package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func main() {
	student1 := Student{
		Name:  "NVA",
		Score: 90,
	}
	student2 := Student{
		Name:  "NVB",
		Score: 92,
	}

	student3 := Student{
		Name:  "NVC",
		Score: 100,
	}

	fmt.Println("Info Student")
	fmt.Println(student1.Name, " - Score: ", student1.Score)
	fmt.Println(student2.Name, " - Score: ", student2.Score)
	fmt.Println(student3.Name, " - Score: ", student3.Score)

	students := []Student{student1, student2, student3}
	fmt.Println("\nFor loop:")
	for _, student := range students {
		fmt.Println(student.Name, " - Score: ", student.Score)
	}
}
