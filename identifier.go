package main

import "fmt"

func main() {
	var a string = "noob"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var a1 int
	fmt.Println(a1)

	var a2 bool
	fmt.Println(a2)

	var d = true
	fmt.Println(d)

	var d1 int
	fmt.Println(d1)

	abc := 1
	fmt.Println(abc)

	var x, y int
	var (
		ab  int
		ab1 bool
	)
	var c1, d2 int = 1, 2
	var e, f = 1, "hello"
	fmt.Println(x, y, ab, ab1, c1, d2, e, f)
}
