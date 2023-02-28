package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Println("第一行为true")
	}
	if a || b {
		fmt.Println("第二行为true")
	}
	if !(a && b) {
		fmt.Print("第三行为TRUE")
	}
}
