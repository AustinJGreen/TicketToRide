package game

import "math"

func Pascal(n float64) float64 {
	return math.Floor((math.Sqrt(1 + 8 * n) - 1) / 2) * math.Floor((math.Sqrt(1 + 8 * n) + 1) / 2) / 2
}

func Signal(n, p float64, a int) int {
	return a * int(math.Sin((math.Pi * n) / (2 * p)))
}

func MinInt(a, b int) int {
	switch a < b {
	case true:
		return a
	default:
		return b
	}
}

func MaxInt(a, b int) int {
	switch a > b {
	case true:
		return a
	default:
		return b
	}
}