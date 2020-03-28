package game

import (
	"math"
)

type TrackCount struct {}

func (TrackCount) Cost(Track) int {
	return 1
}

func (TrackCount) Best(h1, h2 int) int {
	return MinInt(h1, h2)
}

func (TrackCount) Less(h1, h2 int) bool {
	return h1 < h2
}

func (TrackCount) InitHeuristicMap() int {
	return math.MaxInt16
}

func (TrackCount) NegativeInfinity() int {
	return math.MinInt16
}