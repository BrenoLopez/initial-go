package main

import (
	"fmt"
	"time"
)

// chan <- bool quer dizer que so pode escrever na goroutine
func numbers(done chan<- bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
	done <- true
}

func words(done chan<- bool) {
	for w := 'a'; w < 'j'; w++ {
		fmt.Printf("%c ", w)
		time.Sleep(time.Millisecond * 230)
	}
	done <- true
}

func numbers2(channel chan<- int) {
	for i := 0; i < 10; i++ {
		channel <- i
		fmt.Printf("Write on channel %d \n", i)
		time.Sleep(time.Millisecond * 150)
	}
	fmt.Println("End write")
	close(channel)
}

func words2(channel chan<- string) {
	for w := 'a'; w < 'j'; w++ {
		channel <- fmt.Sprintf("%c", w)
		fmt.Printf("%c ", w)
		time.Sleep(time.Millisecond * 90)
	}
	close(channel)
}
func main() {
	// goroutines without buffer
	channelToNumbers := make(chan bool)
	channelToWords := make(chan bool)

	go numbers(channelToNumbers)
	go words(channelToWords)
	<-channelToNumbers
	<-channelToWords

	fmt.Println()
	fmt.Println("End execution")
	// goroutines without buffer
	channelToNumbers2 := make(chan int)
	go numbers2(channelToNumbers2)
	for item := range channelToNumbers2 {
		fmt.Printf("read by the channel %d \n", item)
		time.Sleep(time.Millisecond * 150)

	}
	// goroutines with buffer
	channelToWords2 := make(chan string, 3)
	go words2(channelToWords2)
	for item := range channelToWords2 {
		fmt.Printf("read by the channel %s \n", item)
		time.Sleep(time.Millisecond * 150)

	}
}
