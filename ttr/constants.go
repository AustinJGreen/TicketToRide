package ttr

import (
	"errors"
	"fmt"
	"strings"
)

const (
	Wild = iota
	White
	Black
	Red
	Orange
	Yellow
	Green
	Blue
	Pink
)

var ConnectionScores = []int { 0, 1, 2, 4, 7, 10, 15, 18, 21, 27 }

func ColorToString(color int) (string, error) {
	switch color {
	case Wild:
		return "Wild", nil
	case White:
		return "White", nil
	case Black:
		return "Black", nil
	case Red:
		return "Red", nil
	case Orange:
		return "Orange", nil
	case Yellow:
		return "Yellow", nil
	case Green:
		return "Green", nil
	case Blue:
		return "Blue", nil
	case Pink:
		return "Pink", nil
	default:
		return "", errors.New(fmt.Sprintf("invalid color %d", color))
	}
}

func StringToColor(color string) (int, error) {
	switch strings.ToLower(color) {
	case "wild", "locomotive":
		return Wild, nil
	case "white", "passenger":
		return White, nil
	case "black", "hopper":
		return Black, nil
	case "red", "coal":
		return Red, nil
	case "orange", "freight":
		return Orange, nil
	case "yellow", "reefer":
		return Yellow, nil
	case "green", "caboose":
		return Green, nil
	case "blue", "tanker":
		return Blue, nil
	case "pink", "box":
		return Pink, nil
	default:
		return Wild, errors.New(fmt.Sprintf("invalid color %s", color))
	}
}