package main

type Agent interface {
	Id() int
	Start() string
	Destination() string
	LeaveTime() float32
}

type agent struct {
	ID int
	DEPART_TIME float32

	START string
	DEST  string

	ENVIR int
	CST   int
	TIME  int
}

func (a agent) Id() int {
	return a.ID
}

func (a agent) Start() string {
	return a.START
}

func (a agent) Destination() string {
	return a.DEST
}

func (a agent) LeaveTime() float32 {
	return a.DEPART_TIME
}
