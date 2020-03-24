package ttr

type Heuristic interface {
	Cost(Track) int
	Best(h1, h2 int) int
	Less(h1, h2 int) bool
	InitHeuristicMap() int
	NegativeInfinity() int
}
