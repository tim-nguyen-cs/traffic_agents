package main

type Edge interface {
	Weight() float64
	To() string
	From() string
	Time() int
	AddAgent()
	RemoveAgent()
}

type edge struct {
	weight   float64
	from     string
	to       string
	time     int
	capacity int
}

func (e edge) Weight() float64 {
	return e.weight
}

func (e edge) To() string {
	return e.to
}

func (e edge) From() string {
	return e.from
}

func (e edge) Time() int {
	return e.time
}

func (e edge) AddAgent() {
	e.capacity = e.capacity + 1
}

func (e edge) RemoveAgent() {
	e.capacity = e.capacity - 1
}
