package ttr

type Ticket struct {
	cityNameA string
	cityNameB string
	points    int
}

func NewTicket(nameA, nameB string, points int) *Ticket {
	return &Ticket{nameA, nameB, points}
}

func (t Ticket) GetTargets() []string {
	return []string { t.cityNameA, t.cityNameB }
}