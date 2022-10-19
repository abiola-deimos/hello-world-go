package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")
	fmt.Println()
	for {
		time.Sleep(time.Second * 3)
		fmt.Printf("the time in UTC is now %v", time.Now().UTC())
		fmt.Println()
	}
}
