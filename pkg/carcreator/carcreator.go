package carcreator

import "fmt"

// Car type
type Car struct {
	Model        string
	Color        string
	StartupSound string
	MPG          float32
	FuelTank     float32 //gallons
}

// Start makes StartupSound
func (c Car) Start() {
	fmt.Println(c.StartupSound)
}

// Drive drives the car and adjusts the FuelTank
func (c Car) Drive(miles float32) {
	// c.FuelTank = c.FuelTank - (miles / c.MPG)

	for m := float32(0); m < miles; m++ {
		fmt.Println(c.Model, "driving mile", m)
		c.FuelTank = c.FuelTank - (1 / c.MPG)
		if c.FuelTank < 0 {
			fmt.Println(c.Model, "is out of fuel")
			return
		}
		fmt.Println(c.Model, "new fuel level is", c.FuelTank)

	}
	fmt.Println("After driving", miles, "the fuel level is at", c.FuelTank)

}

// NewCar assigns propoerties
func NewCar(model string, color string) Car {
	c := Car{
		Model:        model,
		Color:        color,
		StartupSound: "Bang",
		MPG:          20.0,
		FuelTank:     12.0,
	}
	if model == "Supra" {

		c.StartupSound = "Vroom"
	}

	if model == "Skyline" {
		c.StartupSound = "Boom"
	}

	return c
}
