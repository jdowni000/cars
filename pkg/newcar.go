package newcar

import "fmt"

// Car type
type Car struct {
	Model        string
	Color        string
	StartupSound string
}

// Start makes StartupSound
func (c Car) Start() {
	fmt.Println(c.StartupSound)
}

func newCar(model string, color string) Car {
	c := Car{
		Model:        model,
		Color:        color,
		StartupSound: "Bang",
	}
	if model == "Supra" {

		c.StartupSound = "Vroom"
	}

	if model == "Skyline" {
		c.StartupSound = "Boom"
	}

	return c
}
