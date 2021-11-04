package main

import (
	"fmt"
)


type Base struct {
	name string
	age  int
}

func (b *Base) AddAge(handler Handler) {
	fmt.Println("base add Age .....")
	handler.Handler()
}

func (b *Base) reName() {
	fmt.Println("rename ing...")
}

type Child struct {
	*Base
	price int
}


func (c *Child) reName() {
	fmt.Println("child rename ing...")
}

type Handler interface{
	Handler()
}

func (c *Child)  Handler(){
	fmt.Println("c handler....")
}

func (b *Base)  Handler(){
	fmt.Println("b handler....")
}

func main() {
	b := &Base{
		name: "base xx",
		age:  100,
	}
	c := &Child{
		Base:  b,
		price: 333,
	}
	c.AddAge(c)
	b.AddAge(b)
}