package main

// Defines the structs and their methods, for parsing data and using it in the sim.
type FeatureCollection struct {
	Features []Feature
}

type Feature struct {
	Properties PropertyList
}

type PropertyList struct {
	TOOLTIP string
	OBJECTID int
}

func (p *PropertyList) get_TOOLTIP() string {
	return p.TOOLTIP
}

func (p *PropertyList) get_OBJECTID() int {
	return p.OBJECTID
}


