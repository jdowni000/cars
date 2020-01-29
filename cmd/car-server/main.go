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

	m["Nissan"] = carcreator.NewCarWithOutput("Skyline", "blue", resp)
	m["Tesla"] = carcreator.NewCarWithOutput("Roadster", "white", resp)
	m["Toyota"] = carcreator.NewCarWithOutput("Supra", "black", resp)
	m["Chevrolet"] = carcreator.NewCarWithOutput("Corvette", "red", resp)
	m["Dodge"] = carcreator.NewCarWithOutput("Viper", "Silver", resp)

	if "Roadster" != userCar {
		resp.Write([]byte("Dude pick a real car like a Roadster"))
		return
	}

	for k, v := range m {
		if userCar == v.Model {
			resp.Write([]byte("The company that makes your car is " + k + "\n"))
			v.Start()
			v.Drive(200)
		}
	}
}
