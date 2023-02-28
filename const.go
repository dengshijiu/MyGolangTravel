package main

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH int = 12
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Println("面积：", area)
	fmt.Println(a, b, " ", c)
	const (
		i = 1 << iota
		j = 109 << iota
		k
		l
	)
	fmt.Println("i=", i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println("l=", l)
}
