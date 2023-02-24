package main

import (
	"fmt"
	"strconv"
)

func hello(name string) {
	fmt.Println("hello " + name)
}

func sum(a, b int) int {
	return a + b
}
func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}
	total = a + i
	return
}
func main() {
	hello("Breno")

	fmt.Println("sum total:", sum(10, 5))
	total, err := convertAndSum(10, "7")
	fmt.Println("sum with convert:", total, err)
}
