package ttr

import (
	"errors"
	"fmt"
	"math"
)



type Track struct {
	id     int
	cityA  *City
	cityB  *City
	length int
	color  int
}

func NewTrack(id int, cityA, cityB *City, color, length int) *Track {
	return &Track{id, cityA, cityB, length, color}
}

func (track Track) Cost(currentCity *City, heuristicMap map[int]int, h Heuristic) int {
	targetCity, _ := track.Target(currentCity)
	currentCost := heuristicMap[currentCity.id]
	targetCost := heuristicMap[targetCity.id]
	calculatedCost := currentCost - targetCost
	return calculatedCost - h.Cost(track)
}

func (track Track) Matches(t *Track) bool {
	return t.HasCity(track.cityA) && t.HasCity(track.cityB)
}

func (track Track) And(targetTrack *Track) (*City, error) {
	if targetTrack.HasCity(track.cityA) {
		return track.cityA, nil
	} else if targetTrack.HasCity(track.cityB) {
		return track.cityB, nil
	}

	return nil, errors.New("tracks do not intersect")
}

func (track Track) Xor(targetTrack *Track) (*City, error) {
	if !targetTrack.HasCity(track.cityA) {
		return track.cityA, nil
	} else if !targetTrack.HasCity(track.cityB) {
		return track.cityB, nil
	}

	return nil, errors.New("tracks have no exclusive cities")
}

func (track Track) HasCity(city *City) bool {
	switch city.id {
	case track.cityA.id, track.cityB.id:
		return true
	}
	return false
}

func (track Track) Target(sourceCity *City) (*City, error) {
	switch sourceCity.id {
	case track.cityA.id:
		return track.cityB, nil
	case track.cityB.id:
		return track.cityA, nil
	}

	return nil, errors.New(fmt.Sprintf("track does not contain city %s", sourceCity))
}

func (track Track) Length() int {
	return track.length
}

func (track Track) Score() int {
	/*
		https://oeis.org/A306221
		1 = 1 pts
		2 = 2 pts
		3 = 4 pts
		4 = 7 pts
		5 = 10 pts
		6 = 15 pts
		7 = 18 pts
		8 = 21 pts
		9 = 27 pts
		1, 2, 3, 3, 5, 3, 3, 6
		1, 1, 0, 2, -2, 0, 3 (Pascals triangle)
		0, -1, 2, -4, 2, 3
		-1, 3, -5, 6, 1
		4, 8, 11, -5
		4, 3, -16
		-1, -19
	*/

	l := float64(track.length)
	s := int(math.Pow(l, 1.4))
	for i := 4.; i <= l; i++ {
		s += Signal(l, i, int(Pascal(i - 3)))
	}

	return s
}