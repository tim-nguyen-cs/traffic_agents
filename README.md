#   Producing Travel Agents
This program uses Go to produce agents that will be used in a traffic simulation model.
Each agent has a starting and destination location and priority factors associated with temporal urgency, environmental consequences, or monetary cost.

Each agent is first randomly associated with an OBJECTID that correlates with their starting and desination locations from the Los Angeles City GeoHub, found [here](http://geohub.lacity.org/datasets/0372aa1fb42a4e29adb9caadcfb210bb_9), under GeoJSON application programming interfaces.

This OBJECTID is then used to generate the associated TOOLTIP that the agent will traverse to.
Each OBJECTID and priority factor is developed randomly using rand.Intn().

## Usage: AddAgent.go
 The AddAgent.go program incorporates flags to parse command-line arguments, specifically requiring the file name of the input data and the number of agents wanted to be produced. The program then writes to an "agents.json" file in the required formatting. Please ensure to include teh features.go package.
 
 ```
[EXECUTION_CODE] -input (GeoJSON file to parse, string) -num (Number of agents to produce, int)
  
```

## Usage: simulate.go
 The simulate.go program incorporates flags to parse command-line arguments, specifically requiring the file name of the input data. The program then writes to an "agents.json" file in the required formatting.
 
 ```
[EXECUTION_CODE] -input (GeoJSON file to parse, string)
 ```
## Errors
  1) Simulation does not yet create or add Edges
  2) Some intersections are omitted because of JSON formatting concerns
