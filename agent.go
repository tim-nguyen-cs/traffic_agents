package main

import "strconv"

type Agent interface {
	Id() string
	Start() string
	Destination() string
	LeaveTime() int
}

type agent struct {
	ID int
	
	START_ID int 
	DEST_ID int
	DEPART_TIME int	

	START_LOCATION string
	DEST_LOCATION string	

	ENVIR int
	CST int
	TIME int
	
}

func (a agent) Id() string {
	return strconv.Itoa(a.ID)
}

func (a agent) Start() string {
	return a.START_LOCATION
}

func (a agent) Destination() string {
	return a.DEST_LOCATION
}

func (a agent) LeaveTime() int {
	return a.DEPART_TIME
}
