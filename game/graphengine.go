package game

import (
	"math"
)

type GraphEngine struct {
	context *Map
}

func NewGraphEngine(m *Map) *GraphEngine {
	return &GraphEngine{m}
}

func (ge GraphEngine) HeuristicMap(startCity *City, h Heuristic) map[*City]int {

	// Stack of cities for iteration
	cityStack := NewHeuristicHeap()
	cityStack.Push(&HeuristicItem{startCity, 0, 0, h})

	// Variable for our visitedCities tracks
	inStack := make(map[*City]bool, 1)
	inStack[startCity] = true

	heuristicMap := make(map[*City]int, 0)

	// Initialize all other heuristics with infinity
	for _, city := range ge.context.Cities() {
		heuristicMap[city] = h.InitHeuristicMap()
	}

	// Set startCity to 0
	heuristicMap[startCity] = 0

	// Calculate all heuristicMap with dijkstra
	for cityStack.Len() > 0 {
		// Pop current city from stack
		curItem := cityStack.Pop()
		curCity := curItem.(*HeuristicItem).value.(*City)

		// Get neighbors and calculate node heuristicMap
		for _, track := range curCity.Tracks() {

			// If it's not a inStack city, calculate distance
			targetCity, _ := track.Target(curCity)

			// Calculate new potential newCost and compare with existing distance
			existingCost := heuristicMap[targetCity]
			newCost := heuristicMap[curCity] + h.Cost(*track)
			heuristicMap[targetCity] = h.Best(existingCost, newCost)

			// If we haven't visited the target city, add it to heap
			if !inStack[targetCity] {
				cityStack.Push(&HeuristicItem{targetCity, newCost, 0, h})
				inStack[targetCity] = true
			}
		}
	}

	return heuristicMap
}

func (ge GraphEngine) BuildSinglePath(source, target *City, h Heuristic) *Path {

	// Calculate all distances from end city
	distances := ge.HeuristicMap(target, h)

	// Start constructing path
	shortestPath := NewPath(source)

	curCity := source
	for curCity != target {

		bestHeuristic := math.MinInt16
		var bestTrack *Track

		// Get all tracks from current city
		for _, track := range curCity.Tracks() {

			// Get target city
			targetCity, _ := track.Target(curCity)

			// Calculate heuristic to get to next distance given connection length
			currentDistance := distances[curCity]
			targetDistance := distances[targetCity]
			heuristic := currentDistance - (targetDistance + h.Cost(*track))

			// Store best route
			if heuristic > bestHeuristic {
				bestHeuristic = heuristic
				bestTrack = track
			}
		}

		// Add track to path
		shortestPath.AddTrack(bestTrack)
		curCity, _ = bestTrack.Target(curCity)
	}

	return shortestPath
}

func (ge GraphEngine) BuildSinglePaths(source, target *City, h Heuristic, error int) []*Path {

	// If source or target is bad, return no paths
	if source == nil || target == nil {
		return []*Path{}
	}

	// Calculate heuristic map from end city
	heuristicMap := ge.HeuristicMap(target, h)

	// Initialize tracks
	bestPaths := make([]*Path, 0)

	// Initialize city queue
	pathQueue := make([]*Path, 1)
	pathQueue[0] = NewPath(source)

	for len(pathQueue) > 0 {
		// Dequeue current city
		curPath := pathQueue[0]
		pathQueue = pathQueue[1:]

		curCity, _ := curPath.LastCity()

		// If we're at our target, stop
		if curCity == target {
			bestPaths = append(bestPaths, curPath)
			continue
		}

		// Keep track of best tracks to use next
		bestCost := h.NegativeInfinity()
		bestTracks := make([]*Track, 0)

		// Get all tracks from current city
		for _, track := range curCity.Tracks() {

			// If the track is already in the path, stop
			if !curPath.CanAdd(track) {
				continue
			}

			// Calculate cost to get to next distance given connection length
			actualCost := track.Cost(curCity, heuristicMap, h)
			allowedError := error + curPath.Error(heuristicMap, h)

			// Prune routes that are invalid
			// This is only valid for heuristic costs that are always positive
			if allowedError < 0 {
				continue
			}

			// Store route if equal or better cost
			if h.Less(bestCost, -allowedError) && !h.Less(actualCost, -allowedError) {
				// cost is better, update cost max and
				// set clear existing best tracks
				bestCost = -allowedError
				bestTracks = bestTracks[:0]
				bestTracks = append(bestTracks, track)
			} else if !h.Less(actualCost, bestCost) {
				// append track to best tracks
				bestTracks = append(bestTracks, track)
			}
		}

		for _, track := range bestTracks {
			updatedPath := curPath.Copy()
			updatedPath.AddTrack(track)
			pathQueue = append(pathQueue, updatedPath)
		}
	}

	return bestPaths
}

func (ge GraphEngine) BuildMultiPath(route Route, h Heuristic) *Path {
	m := ge.context
	path := NewPath(m.CityFromName(route.GetTargets()[0]))
	targets := route.GetTargets()
	for i := 0; i < len(targets) - 1; i++ {
		source := m.CityFromName(targets[i])
		target := m.CityFromName(targets[i + 1])
		path.AddPath(ge.BuildSinglePath(source, target, h))
	}

	return path
}

func (ge GraphEngine) BuildBestPath() {

	/**
	Structures/Method needed:
	SearchEngine that can search tracks to see if a ticket is completed


	Main loop pseudocode
	1. Calculate heuristic for tracks that are worth the most for a starting point.
		1a. For each ticket, calculate all shortest paths.
		1b. For each path, loop through all tracks and calculate the value of that route
		as V(track, route) = track.Score + ((track.Length / route.Length) * route.Score)
	2. If N is the maximum number of trains we can use, sort the tracks by worth and place
	on those routes until we run out of trains.
	3. This is the start of our iteration loop, calculate the score of all of our tracks and tickets.
	4. Start the loop of improving, stop until we cannot find an improvement
		4a. Find an improvement.
		4b. Apply improvment.

	Improvement pseudocode
	1. Calculate each track's contribution to our score.
		1a. Each tracks contribution can be calculated as
		isRouteComplete ? route.Score : -route.Score +
		track.Score
	2. For each t
	*/

	return
}