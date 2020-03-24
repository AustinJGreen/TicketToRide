package ttr

type Path struct {
	tracks []*Track
	sourceCity *City
}

func NewPath(sourceCity *City) *Path {
	p := new(Path)
	p.tracks = make([]*Track, 0)
	p.sourceCity = sourceCity
	return p
}

func (p Path) Copy() *Path {
	copy := NewPath(p.sourceCity)
	for _, track := range p.tracks {
		copy.AddTrack(track)
	}
	return copy
}

func (p Path) HeuristicScore (h Heuristic) int {
	sum := 0
	for _, track := range p.tracks {
		sum += h.Cost(*track)
	}

	return sum
}

func (p Path) Error(heuristicMap map[int]int, h Heuristic) int {
	totalError := 0
	cities := p.Cities()
	for i, track := range p.tracks {
		trackError := track.Cost(cities[i], heuristicMap, h)
		totalError += trackError
	}

	return totalError
}

func (p Path) Hash() int64 {
	var hash int64 = 17
	for _, track := range p.tracks {
		hash = hash * 19 + int64(track.id)
	}
	return hash
}

func (p *Path) AddTrack(track *Track) {
	p.tracks = append(p.tracks, track)
}

func (p Path) CanAdd(t *Track) bool {
	for _, track := range p.tracks {
		if t.id == track.id || track.Matches(t) {
			return false
		}
	}

	return true
}

func (p *Path) AddPath(path *Path) {
	for _, track := range path.tracks {
		p.AddTrack(track)
	}
}

func (p Path) LastTrack() *Track {
	if p.TrackCount() == 0 {
		return nil
	}

	return p.tracks[p.TrackCount() - 1]
}

func (p Path) LastCity() (*City, error) {
	if p.TrackCount() == 0 {
		return p.sourceCity, nil
	} else if p.TrackCount() == 1 {
		return p.tracks[0].Target(p.sourceCity)
	} else {
		return p.tracks[p.TrackCount() - 1].Xor(p.tracks[p.TrackCount() - 2])
	}
}

func (p Path) TrackCount() int {
	return len(p.tracks)
}

func (p Path) TotalLength() int {
	sum := 0
	for i := 0; i < len(p.tracks); i++ {
		sum += p.tracks[i].length
	}

	return sum
}

func (p Path) Tracks() []*Track {
	return p.tracks
}

func (p Path) Cities() []*City {
	cnt := p.TrackCount()
	if cnt == 0 {
		return []*City { }
	} else if cnt == 1 {
		targetCity, _ := p.tracks[0].Target(p.sourceCity)
		return []*City { p.sourceCity, targetCity }
	}

	cities := make([]*City, cnt + 1)
	cities[0], _ = p.tracks[0].Xor(p.tracks[1])
	for i := 1; i < cnt; i++ {
		cities[i], _ = p.tracks[i].And(p.tracks[i - 1])
	}
	cities[cnt], _ = p.tracks[cnt- 1].Xor(p.tracks[cnt- 2])
	return cities
}
