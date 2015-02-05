package main

import (
	"fmt"
)

type Animal interface {
	cry()
}

type Cat struct {
	Age   uint
	Color string
}

func (c Cat) cry() {
	fmt.Println("Meow")
}

type Dog int

func (d *Dog) cry() {
	fmt.Println("Bowwow")
}

func main() {
	c := Cat{Age: 10, Color: "red"}
	fmt.Println(c.Age)
	fmt.Println(c.Color)
	c.cry()

	d := Dog(1)
	d.cry()
}
