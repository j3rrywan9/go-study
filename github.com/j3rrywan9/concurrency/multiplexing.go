package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator: function that returns a channel
func boring(msg string) <-chan string {	// Returns receive-only channel of strings.
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() { for { ch <- <-input1 }}()
	go func() { for { ch <- <-input2 }} ()
	return ch
}

func main() {
	ch := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("You're both boring; I'm leaving.")
}
