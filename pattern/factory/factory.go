package main

import "fmt"

type Product interface {
	Use()
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	fmt.Println("Using Product A")
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	fmt.Println("Using Product B")
}

type Creator interface {
	CreateProduct() Product
}

type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	var creator Creator
	creator = &ConcreteCreatorA{}
	productA := creator.CreateProduct()
	productA.Use()
	creator = &ConcreteCreatorB{}
	productB := creator.CreateProduct()
	productB.Use()
}
