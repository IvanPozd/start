package datas

import "fmt"

func Arrays() {
	var todoList = [4]string{
		"Train with weight jacket",
		"Fuck with wife",
		"Build a house",
	}
	todoList[3] = "Learning Golang"

	fmt.Printf("Count of a list items %d\n", len(todoList))

	for index, item := range todoList {
		newIndex := index + 1
		fmt.Printf("%d: %s\n", newIndex, item)
	}
}

func Cycles() {
	// infinity cycle
	i := 0
	for {
		if i == 100 {
			break
		}
		fmt.Println(i)
		i++
	}
}
