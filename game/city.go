package game

type City struct {
	name   string
	tracks []*Track
}

func CityNew(name string) *City {
	return &City{name, nil}
}

func (city *City) AddPath(path *Track) {
	city.tracks = append(city.tracks, path)
}

func (city City) Tracks() []*Track {
	return city.tracks
}