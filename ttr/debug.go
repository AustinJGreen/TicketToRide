package ttr

func IdMapToReadable(m Map, idMap map[int]int) map[string]int {
	readable := make(map[string]int, len(idMap))
	for key, value := range idMap {
		readable[m.CityFromId(key).name] = value
	}

	return readable
}

func PretendUse(x interface{}) {}
