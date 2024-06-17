package main

import "fmt"

func main() {
	fmt.Println("if else")
	logincount := 23
	var result string

	if logincount < 10 {
		result = "fail"
	} else {
		result = "pass"
	}
	fmt.Println(result)
}
