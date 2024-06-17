package main

import "fmt"

func main() {
	fmt.Println("Slice tutorials")

	var fruitlist = []string{"Apple", "Banana", "Cherry"}
	fmt.Println("Original:", fruitlist)
	fruitlist = append(fruitlist, "Mango")
	fmt.Println("After append:", fruitlist)
	fruitlist = append(fruitlist[1:])
	fmt.Println(fruitlist)
}
