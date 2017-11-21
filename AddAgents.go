package main

//Imports required packages to run program
import (
	"fmt"
	"math/rand"
	"time"

	"encoding/json"
	"flag"
	"log"
	"os"
)

func main() {
	//Receives "input" and "num" command-line arguments
	inputFileName := flag.String("input", "", "Data to parse")
	NumAgents := flag.Int("num", 0, "Number of agents in simulation")
	flag.Parse()

	//Quits if inputFileName or NumAgents are empty
	if *inputFileName == "" || *NumAgents == 0 {
		flag.Usage()
		os.Exit(1)
	}

	//Opens the file
	inputFile, err := os.Open(*inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	//Returns a decoder to parse GeoJSON file
	dec := json.NewDecoder(inputFile)
	var places FeatureCollection
	err = dec.Decode(&places)
	if err != nil {
		log.Fatal(err)
	}

	// Generate TOOLTIP via OBJECTID via associative data type
	ID_TO_TOOLTIP := make(map[int]string)
	for _, feature := range places.Features {
		ID_TO_TOOLTIP[feature.Properties.OBJECTID] = feature.Properties.TOOLTIP
	}

	//De-seeds randomization to produce new values each iteration
	rand.Seed(time.Now().UnixNano())

	//Prints formatted String
	fmt.Println("{")
	fmt.Println("\t\"agents\": [")

	for i := 1; i <= *NumAgents; i++ {

		//Produces random start/destination locations and priority rankings
		START_ID := random_int(2001, 178154)
		DEST_ID := random_int(2001, 178154)

		ENVIR := random_int(1, 10)
		CST := random_int(1, 10)
		TIME := random_int(1, 10)

		DEPART_TIME := rand.Float32() * 3

		//Produces new start/destination locations if START_ID == DEST_ID
		for START_ID == DEST_ID {
			START_ID = random_int(2001, 178154)
			DEST_ID = random_int(2001, 178154)
		}

		//Produces new start/destination locations if associated TOOLTIP does not exist
		for ID_TO_TOOLTIP[START_ID] == "" || ID_TO_TOOLTIP[DEST_ID] == "" {
			START_ID = random_int(2001, 178154)
			DEST_ID = random_int(2001, 178154)
		}

		//Prints result
		fmt.Printf("\t\t{ \"ID\": %v,  \"START\": \"%v\", ", i, ID_TO_TOOLTIP[START_ID])
		fmt.Printf("\"DEST\": \"%v\", \"DEPART_TIME\": %v, ", ID_TO_TOOLTIP[DEST_ID], DEPART_TIME)
		fmt.Printf(" \"ENVIR\": %v, \"CST\": %v, \"TIME\": %v }, \n", ENVIR, CST, TIME)
	}

	fmt.Println("\t]")
	fmt.Println("}")

}

//Produces a random number from [min,max]
func random_int(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
