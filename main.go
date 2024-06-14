package main

import (
	"transport/route"
	"transport/transport"
)

func main() {
    bus := &transport.Bus{Name: "Автобус №1"}
    train := &transport.Train{Name: "Потяг Інтерсіті"}
    plane := &transport.Plane{Name: "Літак МАУ"}

    myRoute := &route.Route{}
    myRoute.AddVehicle(bus)
    myRoute.AddVehicle(train)
    myRoute.AddVehicle(plane)

    myRoute.ShowVehicles()
    myRoute.Travel("Пасажир Іван")
}
