package ttr

import (
	"errors"
	"fmt"
	"math"
)

type Map struct {
	cities []*City
	citiesByName  map[string]*City
	citiesById    map[int]*City
	paths   []*Track
	tickets []*Ticket
	name    string
}

func NewMap() *Map {
	m := new(Map)
	m.cities = make([]*City, 0)
	m.citiesByName = make(map[string]*City, 0)
	m.citiesById = make(map[int]*City, 0)
	m.paths = make([]*Track, 0)
	m.tickets = make([]*Ticket, 0)
	return m
}

func (m Map) CityFromName(name string) *City {
	return m.citiesByName[name]
}

func (m Map) CityFromId(id int) *City {
	return m.citiesById[id]
}

func (m Map) CityNameList() []string {
	names := make([]string, 0, len(m.cities))
	for _, v := range m.citiesByName {
		names = append(names, v.name)
	}
	return names
}

func (m Map) CityTracks(name string) []*Track {
	if c := m.citiesByName[name]; c != nil {
		return c.tracks
	}

	return nil
}

func (m *Map) AddCity(name string) {
	newCity := CityNew(len(m.cities), name)
	m.cities = append(m.cities, newCity)
	m.citiesById[newCity.id] = newCity
	m.citiesByName[newCity.name] = newCity
}

func (m Map) HasCity(name string) bool {
	return m.citiesByName[name] != nil
}

func (m *Map) ConnectCities(nameA, nameB string, color int, length int) {
	if !m.HasCity(nameA) {
		panic(fmt.Sprintf("%s does not contain a city called %s", m.name, nameA))
	}

	if !m.HasCity(nameB) {
		panic(fmt.Sprintf("%s does not contain a city called %s", m.name, nameB))
	}

	cityA := m.citiesByName[nameA]
	cityB := m.citiesByName[nameB]

	// Create 1 track object between 2 cities
	path := NewTrack(len(m.paths), cityA, cityB, color, length)

	// Append to list of paths
	m.paths = append(m.paths, path)

	// Add path to each city
	cityA.AddPath(path)
	cityB.AddPath(path)
}

func (m *Map) AddTicket(nameA, nameB string, points int) error {
	if !m.HasCity(nameA) {
		return errors.New(fmt.Sprintf("%s does not contain a city called %s", m.name, nameA))
	}

	if !m.HasCity(nameB) {
		return errors.New(fmt.Sprintf("%s does not contain a city called %s", m.name, nameB))
	}

	ticket := NewTicket(nameA, nameB, points)
	m.tickets = append(m.tickets, ticket)
	return nil
}

func (m Map) HeuristicMap(startCity *City, h Heuristic) map[int]int {

	// Stack of cities for iteration
	cityStack := NewHeuristicHeap()
	cityStack.Push(&HeuristicItem{startCity, 0, 0, h})

	// Variable for our visitedCities tracks
	inStack := make(map[int]bool, 1)
	inStack[startCity.id] = true

	heuristicMap := make(map[int]int, 0)

	// Initialize all other heuristics with infinity
	for _, city := range m.cities {
		heuristicMap[city.id] = h.InitHeuristicMap()
	}

	// Set startCity to 0
	heuristicMap[startCity.id] = 0

	// Calculate all heuristicMap with dijkstra
	for cityStack.Len() > 0 {
		// Pop current city from stack
		curItem := cityStack.Pop()
		curCity := curItem.(*HeuristicItem).value.(*City)

		// Get neighbors and calculate node heuristicMap
		for _, track := range curCity.tracks {

			// If it's not a inStack city, calculate distance
			targetCity, _ := track.Target(curCity)

			// Calculate new potential newCost and compare with existing distance
			existingCost := heuristicMap[targetCity.id]
			newCost := heuristicMap[curCity.id] + h.Cost(*track)
			heuristicMap[targetCity.id] = h.Best(existingCost, newCost)

			// If we haven't visited the target city, add it to heap
			if !inStack[targetCity.id] {
				cityStack.Push(&HeuristicItem{targetCity, newCost, 0, h})
				inStack[targetCity.id] = true
			}
		}
	}

	return heuristicMap
}

func (m Map) BuildSinglePath(source, target *City, h Heuristic) *Path {

	// Calculate all distances from end city
	distances := m.HeuristicMap(target, h)

	// Start constructing path
	shortestPath := NewPath(source)

	curCity := source
	for curCity.id != target.id {

		bestHeuristic := math.MinInt16
		var bestTrack *Track

		// Get all tracks from current city
		for _, track := range m.CityTracks(curCity.name) {

			// Get target city
			targetCity, _ := track.Target(curCity)

			// Calculate heuristic to get to next distance given connection length
			currentDistance := distances[curCity.id]
			targetDistance := distances[targetCity.id]
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

func (m Map) BuildSinglePaths(source, target *City, h Heuristic, error int) []*Path {
	// Calculate heuristic map from end city
	heuristicMap := m.HeuristicMap(target, h)
	readable := IdMapToReadable(m, heuristicMap)
	PretendUse(readable)

	// Initialize paths
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
		if curCity.id == target.id {
			bestPaths = append(bestPaths, curPath)
			continue
		}

		// Keep track of best tracks to use next
		bestCost := h.NegativeInfinity()
		bestTracks := make([]*Track, 0)

		// Get all tracks from current city
		for _, track := range curCity.tracks {

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

func (m Map) BuildMultiPath(route Route, h Heuristic) *Path {
	path := NewPath(m.citiesByName[route.GetTargets()[0]])
	targets := route.GetTargets()
	for i := 0; i < len(targets) - 1; i++ {
		source := m.citiesByName[targets[i]]
		target := m.citiesByName[targets[i + 1]]
		path.AddPath(m.BuildSinglePath(source, target, h))
	}

	return path
}