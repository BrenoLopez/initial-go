package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func words() {
	for w := 'a'; w < 'j'; w++ {
		fmt.Printf("%c ", w)
		time.Sleep(time.Millisecond * 230)
	}
}
func main() {
	go numbers()
	fmt.Println()
	go words()
	go time.Sleep(5 * time.Second)
	fmt.Println()
	fmt.Println("End execution")
}
