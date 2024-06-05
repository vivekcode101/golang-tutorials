package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")
	var ptr1 *int
	fmt.Println("This is the pointer", ptr1)

	myNumber := 23
	var ptr2 = &myNumber
	fmt.Println("This is the pointer", ptr2)
	fmt.Println("This is the value of pointer", *ptr2)
}
