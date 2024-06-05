package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array segment")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[2] = "Cherry"
	fruitlist[3] = "Date"

	fmt.Println(fruitlist)
	fmt.Println(len(fruitlist))

	var veglist = [3]string{"spinach", "cabbage", "Beans"}
	fmt.Println(veglist)
}
