package main

import (
	"fmt"

	"github.com/lrnholanda/go-intensivo-jul/entity"
)

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model, "is started")
}

func (c *Car) ChangeColor(color string) {
	c.Color = color
	fmt.Println("New color: ", c.Color)
}

func main() {
	order, err := entity.NewOrder("156899009", 10, 1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(order.ID)
	}
}
