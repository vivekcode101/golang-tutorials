package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time")
	presenttime := time.Now()
	fmt.Println(presenttime.Format("01-07-2006 15:05:09 Monday"))
}
