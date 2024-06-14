package transport

type Transport interface {
	BoardPassengers(count int) string
	DisembarkPassengers(count int) string
}
