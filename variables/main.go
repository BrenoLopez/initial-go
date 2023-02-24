package main

import "fmt"

// package variables
var (
	name string
	age  int
)

func main() {
	var test, test1 int
	test++
	fmt.Println(test, test1)
	var a, b, c = true, "doidera", 2.4
	fmt.Println(a, b, c)
	name = "Breno Lopes"
	age = 25
	fmt.Println("name:", name)
	fmt.Println("age:", age)

	var x, y = 10, 20
	y, x = x, y
	fmt.Println(x, y)
}
