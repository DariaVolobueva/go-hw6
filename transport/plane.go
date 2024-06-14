package transport

import "fmt"

type Plane struct {
    Name string
}

func (p *Plane) BoardPassengers(count int) string {
    return fmt.Sprintf("%s: Прийнято %d пасажирів", p.Name, count)
}

func (p *Plane) DisembarkPassengers(count int) string {
    return fmt.Sprintf("%s: Висаджено %d пасажирів", p.Name, count)
}
