package valueobject

type Sport struct {
	name string
}

func NewSport(name string) Sport {
	return Sport{
		name: name,
	}
}

func SportsFromArray(sports []string) []Sport {
	result := make([]Sport, 0, len(sports))
	for _, sport := range sports {
		result = append(result, NewSport(sport))
	}
	return result
}

func (s Sport) String() string {
	return s.name
}

var (
	Baseball = Sport{"baseball"}
	Football = Sport{"football"}
)
