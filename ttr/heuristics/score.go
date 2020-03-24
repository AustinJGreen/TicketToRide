package heuristics

import (
	"TicketToRide/ttr"
	"math"
)

type Score struct {}

func (Score) Cost(track ttr.Track) int {
	return track.Score()
}

func (Score) Best(h1, h2 int) int {
	return ttr.MaxInt(h1, h2)
}

func (Score) Less(h1, h2 int) bool {
	return h1 > h2
}

func (Score) InitHeuristicMap() int {
	return math.MinInt16
}

func (Score) NegativeInfinity() int {
	return math.MinInt16
}