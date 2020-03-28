package game

import (
	"math"
)

type Distance struct {}

func (Distance) Cost(track Track) int {
	return track.Length()
}

func (Distance) Best(h1, h2 int) int {
	return MinInt(h1, h2)
}

func (Distance) Less(h1, h2 int) bool {
	return h1 < h2
}

func (Distance) InitHeuristicMap() int {
	return math.MaxInt16
}

func (Distance) NegativeInfinity() int {
	return math.MinInt16
}

