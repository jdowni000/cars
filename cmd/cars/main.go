package main

import (
	"fmt"

	"github.com/jdowni000/cars/pkg/carcreator"
)

// main makes cars with info
func main() {

	m := make(map[string]carcreator.Car)

	m["Nissan"] = carcreator.NewCar("Skyline", "blue")
	m["Tesla"] = carcreator.NewCar("Roadster", "white")
	m["Toyota"] = carcreator.NewCar("Supra", "black")
	m["Chevrolet"] = carcreator.NewCar("Corvette", "red")
	m["Dodge"] = carcreator.NewCar("Viper", "Silver")

	for k, v := range m {
		// fmt.Println(k, v.Model)

		if k == "Nissan" {
			fmt.Println(v.Model)
			v.Start()
			v.Drive(200)
		} else {
			v.Start()
			v.Drive(400)

		}
	}
}
