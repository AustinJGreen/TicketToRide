package ttr

type City struct {
	id     int
	name   string
	tracks []*Track
}

func CityNew(id int, name string) *City {
	return &City{id, name, nil}
}

func (city *City) AddPath(path *Track) {
	city.tracks = append(city.tracks, path)
}