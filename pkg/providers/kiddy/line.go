package kiddy

import (
	"encoding/json"
	"strings"
)

type Line struct {
	sport string
	rate  string
}

func (l Line) Sport() string {
	return l.sport
}

func (l Line) Rate() string {
	return l.rate
}

// UnmarshalJSON {"SPORT": "0.123"}
func (l *Line) UnmarshalJSON(data []byte) error {
	var v map[string]string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	for sportRaw, rateRaw := range v {
		l.sport = strings.ToLower(sportRaw)
		l.rate = rateRaw
	}

	return nil
}
