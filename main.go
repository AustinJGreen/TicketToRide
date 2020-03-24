package main

import (
	"TicketToRide/ttr"
	"TicketToRide/ttr/heuristics"
	"fmt"
)

func main() {

	// Montréal
	// 110 train car cards
	// 12 of each color
	// 14 wilds
	// 30 destination cards
	america, err := ttr.ReadMap("america.map")
	if err != nil {
		fmt.Printf("Failed to load map, %s", err.Error())
		return
	}

	//route := ttr.NewRoute("Nashville", "Sault St. Marie")
	from := america.CityFromName("Los Angeles")
	to := america.CityFromName("Montréal")
	heuristic := heuristics.Distance{}
	paths := america.BuildSinglePaths(from, to, heuristic, 0)

	for i, path := range paths {
		fmt.Printf("--- Path %d %d ---\n", i, path.Hash())
		for _, city := range path.Cities() {
			fmt.Println(city)
		}
		fmt.Printf("Score = %d\n", path.HeuristicScore(heuristic))
		fmt.Printf("Length = %d\n", path.TotalLength())
		fmt.Println("--- End Path ---")
	}

	fmt.Printf("%d total paths", len(paths))
	return
}
