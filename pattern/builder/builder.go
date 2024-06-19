package main

import "fmt"

type Car struct {
	Brand        string
	Model        string
	Year         int
	Color        string
	Transmission string
}

type CarBuilder interface {
	SetBrand(brand string) CarBuilder
	SetModel(model string) CarBuilder
	SetYear(year int) CarBuilder
	SetColor(color string) CarBuilder
	SetTransmission(transmission string) CarBuilder
	Build() *Car
}

type ConcreteCarBuilder struct {
	brand        string
	model        string
	year         int
	color        string
	transmission string
}

func NewConcreteCarBuilder() *ConcreteCarBuilder {
	return &ConcreteCarBuilder{}
}

func (b *ConcreteCarBuilder) SetBrand(brand string) CarBuilder {
	b.brand = brand
	return b
}

func (b *ConcreteCarBuilder) SetModel(model string) CarBuilder {
	b.model = model
	return b
}

func (b *ConcreteCarBuilder) SetYear(year int) CarBuilder {
	b.year = year
	return b
}

func (b *ConcreteCarBuilder) SetColor(color string) CarBuilder {
	b.color = color
	return b
}

func (b *ConcreteCarBuilder) SetTransmission(transmission string) CarBuilder {
	b.transmission = transmission
	return b
}

func (b *ConcreteCarBuilder) Build() *Car {
	return &Car{
		Brand:        b.brand,
		Model:        b.model,
		Year:         b.year,
		Color:        b.color,
		Transmission: b.transmission,
	}
}

type Director struct {
	builder CarBuilder
}

func NewDirector(builder CarBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct(brand, model string, year int, color, transmission string) *Car {
	return d.builder.SetBrand(brand).
		SetModel(model).
		SetYear(year).
		SetColor(color).
		SetTransmission(transmission).
		Build()
}

func main() {
	builder := NewConcreteCarBuilder()
	director := NewDirector(builder)
	car := director.Construct("Toyota", "Camry", 2022, "Blue", "Automatic")
	fmt.Printf("Car: %+v\n", *car)
}
