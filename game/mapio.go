package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//
func ReadMap(filename string) (*Map, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	m := NewMap()
	scanner := bufio.NewScanner(file)

	for curLine := 0; scanner.Scan(); curLine++ {
		line := scanner.Text()
		arguments := strings.Split(line, " ")
		if len(arguments) < 1 {
			return nil, errors.New(fmt.Sprintf("invalid map token (line %d)", curLine))
		}

		switch token := arguments[0]; token {
		case "map":
			m.name = strings.Join(arguments[1:], " ")
		case "city":
			cityName := strings.Join(arguments[1:], " ")
			m.AddCity(cityName)
		case "track":
			if len(arguments) < 6 {
				return nil, errors.New(fmt.Sprintf("missing track tokens (line %d)", curLine))
			}

			verb := arguments[1]
			length, err := strconv.Atoi(arguments[len(arguments) - 1])
			if err != nil {
				return nil, errors.New(fmt.Sprintf("invalid track length (line %d)", curLine))
			}

			color := arguments[len(arguments) - 2]
			cities := strings.Split(strings.Join(arguments[2:len(arguments) - 2], " "), " and ")
			if len(cities) != 2 {
				return nil, errors.New(fmt.Sprintf("track should contain two cities (line %d)", curLine))
			}

			switch verb {
			case "both":
				color, err := StringToColor(color)
				if err != nil {
					return nil, err
				}

				m.ConnectCities(cities[0], cities[1], color, length)
			}
		case "ticket":
			cities := strings.Split(strings.Join(arguments[1:len(arguments) - 1], " "), " and ")
			if len(cities) != 2 {
				return nil, errors.New(fmt.Sprintf("missing ticket tokens (line %d)", curLine))
			}

			points, err := strconv.Atoi(arguments[len(arguments) - 1])
			if err != nil {
				return nil, errors.New(fmt.Sprintf("invalid ticket score (line %d)", curLine))
			}

			if err := m.AddTicket(cities[0], cities[1], points); err != nil {
				return nil, err
			}
		case "test":
			verb := arguments[1]
			switch verb {
			case "tracks":
				city := strings.Join(arguments[2:len(arguments) - 1], " ")
				count, err := strconv.Atoi(arguments[len(arguments) - 1])
				if err != nil {
					return nil, errors.New(fmt.Sprintf("invalid path count (line %d)", curLine))
				}

				if actualCnt := len(m.CityTracks(city)); actualCnt != count {
					return nil, errors.New(fmt.Sprintf("%s contains %d tracks (line %d)", city, actualCnt, curLine))
				}
			case "cities":
				count, err := strconv.Atoi(arguments[len(arguments) - 1])
				if err != nil {
					return nil, errors.New(fmt.Sprintf("invalid city count (line %d)", curLine))
				}

				if len(m.cities) != count {
					return nil, errors.New(fmt.Sprintf("%s doesn't contain %d cities (line %d)", m.name, count, curLine))
				}
			}

		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
