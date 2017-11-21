package main

import	(
	"encoding/json"
	"log"
	"os"
	"flag"
	"fmt"
	//tf "github.com/spencer-p/traffic"
)


/*
=========================
SIMULATION SECTION
=========================
*/

type Simulation struct	{
	AgentArray []Agent
	EdgeArray []Edge
}

func NewSimulation() Simulation	{
	return Simulation{make([]Agent, 0), make([]Edge, 0)}
}


func (s *Simulation) AddAgent(a Agent)	{
	s.AgentArray = append(s.AgentArray, a)
}

func (s *Simulation) AddEdge(e Edge)	{
	s.EdgeArray = append(s.EdgeArray, e)
}

func (s *Simulation) Simulate()	{
	fmt.Println("Simulation in progress")
	//Insert simulation code here.
}



/*
=========================
MAIN FUNCTION
=========================
*/

func main() {

	//Receives "input" and "num" command-line arguments
	inputFileName := flag.String("input", "", "GeoData to parse")
	flag.Parse()

	//Quits if inputFileName is empty
	if *inputFileName == "" {
		flag.Usage()
		os.Exit(1)
	}


	//Opens data file
	inputFile, err := os.Open(*inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	//Returns decoder to parse GeoJSON file
	data_dec := json.NewDecoder(inputFile)
	var places FeatureCollection
	err = data_dec.Decode(&places)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Parsed", len(places.Features), "streets")	


	//Opens "agents.json" file
	input, err:= os.Open("agents.json")
	if err != nil {
		log.Fatal(err)
	}
	//Returns a decoder to prase file of agents
	agent_dec := json.NewDecoder(input)
	var people struct{ Agents []agent }
	err = agent_dec.Decode(&people)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Parsed", len(people.Agents), "agents")
	

	//Makes a new simulation
	newSimulation := NewSimulation()

	
/*	for i, _ := range places.Features {
		newSimulation.AddEdge(places.Features[i])
	}

*/
	//Adds agents to simulation
	for i, _ := range people.Agents {
		newSimulation.AddAgent(people.Agents[i])
	}
	
	log.Println("Added", len(people.Agents), "agents to the simulation")
	
	//Runs simulation
	newSimulation.Simulate()
}
