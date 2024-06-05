package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input"
	fmt.Println(welcome)

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	input, _ := reader1.ReadString('\n')
	fmt.Println("Thanks for your name", input)
}
