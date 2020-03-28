package game

import (
	"errors"
	"fmt"
)

type Map struct {
	cities       []*City
	citiesByName map[string]*City
	tracks       []*Track
	tickets      []*Ticket
	name         string
}

func NewMap() *Map {
	m := new(Map)
	m.cities = make([]*City, 0)
	m.citiesByName = make(map[string]*City, 0)
	m.tracks = make([]*Track, 0)
	m.tickets = make([]*Ticket, 0)
	return m
}

func (m Map) CityFromName(name string) *City {
	return m.citiesByName[name]
}

func (m Map) Cities() []*City {
	return m.cities
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
	newCity := CityNew(name)
	m.cities = append(m.cities, newCity)
	m.citiesByName[newCity.name] = newCity
}

func (m Map) HasCity(name string) bool {
	return m.citiesByName[name] != nil
}

func (m *Map) ConnectCities(nameA, nameB string, color int, length int) error {
	if !m.HasCity(nameA) {
		return errors.New(fmt.Sprintf("%s does not contain a city called %s", m.name, nameA))
	}

	if !m.HasCity(nameB) {
		return errors.New(fmt.Sprintf("%s does not contain a city called %s", m.name, nameB))
	}

	cityA := m.citiesByName[nameA]
	cityB := m.citiesByName[nameB]

	// Create 1 track object between 2 cities
	path := NewTrack(cityA, cityB, color, length)

	// Append to list of tracks
	m.tracks = append(m.tracks, path)

	// Add path to each city
	cityA.AddPath(path)
	cityB.AddPath(path)
	return nil
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

