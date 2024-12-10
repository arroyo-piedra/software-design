package builder

import "fmt"

// Final product that we want to build.
type Car struct {
	Brand    string
	Model    string
	Engine   string
	Color    string
	NumDoors int
}

// CarBuilder is the builder that constructs the Car step by step.
type CarBuilder struct {
	brand    string
	model    string
	engine   string
	color    string
	numDoors int
}

// NewCarBuilder creates a new Car.
func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

// SetBrand set the brand of the car.
func (cb *CarBuilder) SetBrand(brand string) *CarBuilder {
	cb.brand = brand
	return cb
}

// SetModel set the model of the car.
func (cb *CarBuilder) SetModel(model string) *CarBuilder {
	cb.model = model
	return cb
}

// SetEngine set the engine of the car.
func (cb *CarBuilder) SetEngine(engine string) *CarBuilder {
	cb.engine = engine
	return cb
}

// SetColor set the color of the car.
func (cb *CarBuilder) SetColor(color string) *CarBuilder {
	cb.color = color
	return cb
}

// SetNumDoors set the number of doors of the car.
func (cb *CarBuilder) SetNumDoors(numDoors int) *CarBuilder {
	cb.numDoors = numDoors
	return cb
}

// Build constructs the Car.
func (cb *CarBuilder) Build() Car {
	car := Car{
		Brand:    cb.brand,
		Model:    cb.model,
		Engine:   cb.engine,
		Color:    cb.color,
		NumDoors: cb.numDoors,
	}
	fmt.Printf("Car: %+v\n", car)
	return car
}

// This is a struct to be invoked by main.go
type BuilderPattern struct{}

func (b BuilderPattern) ProcessPattern() bool {
	fmt.Println("------------ Builder Pattern ------------")

	toyota := NewCarBuilder().
		SetBrand("Toyota").
		SetModel("Corolla").
		SetEngine("Hybrid").
		SetColor("Red").
		SetNumDoors(4).
		Build()

	fmt.Printf("%s Created\n", toyota.Brand)

	suzuki := NewCarBuilder().
		SetBrand("Suzuki").
		SetModel("Fronx").
		SetEngine("Hybrid").
		SetColor("Black").
		SetNumDoors(4).
		Build()

	fmt.Printf("%s Created\n", suzuki.Brand)

	return true
}
