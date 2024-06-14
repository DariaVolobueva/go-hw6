package route

import (
	"fmt"
	"transport/transport"
)

type Route struct {
    vehicles []transport.Transport
}

func (r *Route) AddVehicle(vehicle transport.Transport) {
    r.vehicles = append(r.vehicles, vehicle)
}

func (r *Route) ShowVehicles() {
    fmt.Println("Маршрут включає наступні транспортні засоби:")
    for _, vehicle := range r.vehicles {
        fmt.Printf("%T\n", vehicle)
    }
}

func (r *Route) Travel(passenger string) {
    fmt.Printf("%s розпочинає подорож.\n", passenger)
    for _, vehicle := range r.vehicles {
        fmt.Println(vehicle.BoardPassengers(1))
        fmt.Println(vehicle.DisembarkPassengers(1))
    }
    fmt.Printf("%s завершив подорож.\n", passenger)
}
