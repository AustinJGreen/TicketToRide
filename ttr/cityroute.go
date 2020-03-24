package ttr

type CityRoute []string

func NewRoute(targets ...string) CityRoute {
	return targets
}

func (cr CityRoute) GetTargets() []string {
	return cr
}