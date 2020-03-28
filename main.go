package main

import (
	"TicketToRide/game"
	"fmt"
)

func main() {

	// Montréal
	// 110 train car cards
	// 12 of each color
	// 14 wilds
	// 30 destination cards
	america, err := game.ReadMap("america.map")
	if err != nil {
		fmt.Printf("Failed to load map, %s", err.Error())
		return
	}

	ge := game.NewGraphEngine(america)
	heuristic := game.Distance{}

	from := america.CityFromName("Los Angeles")
	to := america.CityFromName("Montréal")
	paths := ge.BuildSinglePaths(to, from, heuristic, 0)

	for i, path := range paths {
		fmt.Printf("--- Path %d ---\n", i)
		for _, city := range path.Cities() {
			fmt.Println(city)
		}
		fmt.Printf("Score = %d\n", path.HeuristicScore(heuristic))
		fmt.Printf("Length = %d\n", path.HeuristicScore(game.Distance{}))
		fmt.Println("--- End Path ---")
	}

	fmt.Printf("%d total paths", len(paths))
	return
}
