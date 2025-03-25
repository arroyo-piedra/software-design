package decorator

import (
	"fmt"
)

// In this example we will create a decorator pattern for a coffee shop.
// In this case we can have a base beverage (Espresso, Latte, Tea) and we can add
// different addons (Milk, Sugar, Whipped Cream) to the beverage.

// Step 1: Component Interface
// Defines the behavior
type Beverage interface {
	Cost() float64
	Description() string
}

// Step 2: Concrete Components (Base Beverages)
// Base beverages with different prices and descriptions
type Espresso struct{}

func (e *Espresso) Cost() float64 {
	return 3.00
}

func (e *Espresso) Description() string {
	return "Espresso"
}

type Latte struct{}

func (l *Latte) Cost() float64 {
	return 4.00
}

func (l *Latte) Description() string {
	return "Latte"
}

type Tea struct{}

func (t *Tea) Cost() float64 {
	return 2.50
}

func (t *Tea) Description() string {
	return "Tea"
}

// Step 3: Base Decorator
// Holds a reference to a Beverage object
type BeverageDecorator struct {
	beverage Beverage
}

func (b *BeverageDecorator) Cost() float64 {
	return b.beverage.Cost()
}

func (b *BeverageDecorator) Description() string {
	return b.beverage.Description()
}

// Step 4: Concrete Decorators
// Dynamically modify the cost and description
type Milk struct {
	BeverageDecorator
}

func (m *Milk) Cost() float64 {
	return m.beverage.Cost() + 0.50
}

func (m *Milk) Description() string {
	return m.beverage.Description() + ", Milk"
}

type Sugar struct {
	BeverageDecorator
}

func (s *Sugar) Cost() float64 {
	return s.beverage.Cost() + 0.25
}

func (s *Sugar) Description() string {
	return s.beverage.Description() + ", Sugar"
}

type WhippedCream struct {
	BeverageDecorator
}

func (w *WhippedCream) Cost() float64 {
	return w.beverage.Cost() + 0.75
}

func (w *WhippedCream) Description() string {
	return w.beverage.Description() + ", Whipped Cream"
}

// This is a struct to be invoked by main.go
type DecoratorPattern struct{}

// Step 5: Testing Different Combinations
// Beverages are wrapped with multiple decorators at runtime
func (d DecoratorPattern) ProcessPattern() bool {
	// Base Beverages
	espresso := &Espresso{}
	latte := &Latte{}
	tea := &Tea{}

	// Customizing Drinks
	espressoWithMilk := &Milk{BeverageDecorator{espresso}}
	latteWithSugarAndCream := &WhippedCream{BeverageDecorator{&Sugar{BeverageDecorator{latte}}}}
	teaWithMilkAndSugar := &Sugar{BeverageDecorator{&Milk{BeverageDecorator{tea}}}}

	// Printing Results
	fmt.Println(espressoWithMilk.Description(), "=> Cost: $", espressoWithMilk.Cost())
	fmt.Println(latteWithSugarAndCream.Description(), "=> Cost: $", latteWithSugarAndCream.Cost())
	fmt.Println(teaWithMilkAndSugar.Description(), "=> Cost: $", teaWithMilkAndSugar.Cost())

	return true
}
