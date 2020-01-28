package carcreator

import (
	"fmt"
	"io"
	"os"
)

// Car type
type Car struct {
	Model        string
	Color        string
	StartupSound string
	MPG          float32
	FuelTank     float32 //gallons
	Output       io.Writer
}

// Start makes StartupSound
func (c Car) Start() {
	fmt.Fprintln(c.Output, c.StartupSound)
}

// Drive drives the car and adjusts the FuelTank
func (c Car) Drive(miles float32) {
	// c.FuelTank = c.FuelTank - (miles / c.MPG)

	for m := float32(0); m < miles; m++ {
		fmt.Fprintln(c.Output, c.Model, "driving mile", m)
		c.FuelTank = c.FuelTank - (1 / c.MPG)
		if c.FuelTank < 0 {
			fmt.Fprintln(c.Output, c.Model, "is out of fuel")
			return
		}
		fmt.Fprintln(c.Output, c.Model, "new fuel level is", c.FuelTank)

	}
	fmt.Fprintln(c.Output, "After driving", miles, "the fuel level is at", c.FuelTank)

}

// NewCar assigns propoerties
func NewCar(model string, color string) Car {
	c := Car{
		Model:        model,
		Color:        color,
		StartupSound: "Bang",
		MPG:          20.0,
		FuelTank:     12.0,
		Output:       os.Stdout,
	}
	if model == "Supra" {

		c.StartupSound = "Vroom"
	}

	if model == "Skyline" {
		c.StartupSound = "Boom"
	}

	return c
}

// NewCarWithOutput creates a Car utilizing NewCar with different output
func NewCarWithOutput(model string, color string, output io.Writer) Car {
	c := NewCar(model, color)
	c.Output = output
	return c
}
