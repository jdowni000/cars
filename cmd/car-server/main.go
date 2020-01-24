package main

import (
	"log"
	"net/http"

	"github.com/jdowni000/cars/pkg/carcreator"
)

func main() {
	http.HandleFunc("/", handleCarRequest)
	log.Println(http.ListenAndServe(":8000", nil))
}

func handleCarRequest(resp http.ResponseWriter, req *http.Request) {
	userCar := req.FormValue("car")
	resp.Write([]byte("Your vehicle is a " + userCar + "\n"))

	m := make(map[string]carcreator.Car)

	m["Nissan"] = carcreator.NewCar("Skyline", "blue")
	m["Tesla"] = carcreator.NewCar("Roadster", "white")
	m["Toyota"] = carcreator.NewCar("Supra", "black")
	m["Chevrolet"] = carcreator.NewCar("Corvette", "red")
	m["Dodge"] = carcreator.NewCar("Viper", "Silver")

	for k, v := range m {
		if v.Model == userCar {
			resp.Write([]byte("The company that makes your car is " + k + "\n"))
			v.Start()
			v.Drive(200)
		} else {
			resp.Write([]byte("Dude, pick a real car like a Skyline, Roadster, Supra, Corvette, or Viper"))
		}

	}

}
