package main

import "fmt"

type Strategy interface {
	Execute(a, b int) int
}

type AddStrategy struct{}

func (s *AddStrategy) Execute(a, b int) int {
	return a + b
}

type MultiplyStrategy struct{}

func (s *MultiplyStrategy) Execute(a, b int) int {
	return a * b
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func main() {
	context := &Context{}
	context.SetStrategy(&AddStrategy{})
	fmt.Println("10 + 5 =", context.ExecuteStrategy(10, 5))
	context.SetStrategy(&MultiplyStrategy{})
	fmt.Println("10 * 5 =", context.ExecuteStrategy(10, 5))
}
