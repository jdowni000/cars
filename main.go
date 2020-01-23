package main

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

func main() {

	m := make(map[string]Car)

	m["Nissan"] = newCar("Skyline", "blue")
	m["Tesla"] = newCar("Roadster", "white")
	m["Toyota"] = newCar("Supra", "black")
	m["Chevrolet"] = newCar("Corvette", "red")
	m["Dodge"] = newCar("Viper", "Silver")

	for _, v := range m {
		// fmt.Println(k, v.Model)

		// if k == "Nissan" {
		fmt.Println(v.Model)
		v.Start()
		// }
	}
}
