package game

import (
	"math"
)

type Score struct {}

func (Score) Cost(track Track) int {
	return track.Score()
}

func (Score) Best(h1, h2 int) int {
	return MaxInt(h1, h2)
}

func (Score) Less(h1, h2 int) bool {
	return h1 > h2
}

func (Score) InitHeuristicMap() int {
	return math.MinInt16
}

func (Score) NegativeInfinity() int {
	return math.MaxInt16
}