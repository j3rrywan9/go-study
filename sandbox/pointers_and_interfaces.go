package main

import (
	"fmt"
)

// Animal interface
type Animal interface {
	Speak() string
}

// Cat struct
type Cat struct {
}

// Speak method of Cat with a pointer receiver
func (c *Cat) Speak() string {
	return "Meow!"
}

// Dog struct
type Dog struct {
}

// Speak method of Dog with a pointer receiver
func (d *Dog) Speak() string {
	return "Woof!"
}

func main() {
	animals := []Animal{new(Cat), new(Dog)}
	// for loop with range
	for _, animal := range animals {
		fmt.Printf(animal.Speak() + "\n")
	}
}
