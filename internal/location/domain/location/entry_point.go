package location

import "time"

type (
	EntryPoint struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lng"`
		CreatedAt time.Time
	}
)
