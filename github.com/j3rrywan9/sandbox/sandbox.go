package main

import (
	"fmt"
)

// interface Animal
type Animal interface {
	Speak() string
}

// type Cat
type Cat struct {
}

// method Speak() of type Cat with a pointer receiver
func (c *Cat) Speak() string {
	return "Meow!"
}

// type Dog
type Dog struct {
}

// method Speak() of type Dog with a pointer receiver
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
