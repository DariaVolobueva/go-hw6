package transport

import "fmt"

type Bus struct {
    Name string
}

func (b *Bus) BoardPassengers(count int) string {
    return fmt.Sprintf("%s: Прийнято %d пасажирів", b.Name, count)
}

func (b *Bus) DisembarkPassengers(count int) string {
    return fmt.Sprintf("%s: Висаджено %d пасажирів", b.Name, count)
}