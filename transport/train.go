package transport

import "fmt"

type Train struct {
    Name string
}

func (t *Train) BoardPassengers(count int) string {
    return fmt.Sprintf("%s: Прийнято %d пасажирів", t.Name, count)
}

func (t *Train) DisembarkPassengers(count int) string {
    return fmt.Sprintf("%s: Висаджено %d пасажирів", t.Name, count)
}
