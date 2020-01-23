package main

import (
	"fmt"

	"github.com/jdowni000/cars/pkg/newcar"
)

// main makes cars with info
func main() {

	m := make(map[string]newcar.Car)

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
