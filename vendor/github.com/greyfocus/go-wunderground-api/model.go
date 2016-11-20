package wunderground_api

import (
	"fmt"
	"time"
)

type Request struct {
	Features []string
	Location string
}

type Response struct {
	CurrentConditions *Conditions
}

// The actual types returned by the API are not very reliable. The parser
// attempts to normalize the fields in order to map the data to some fixed
// types.
//
// Unfortunately this does not work very well for the cases when "N/A" is
// returned - for numeric types, "-1" will be returned and for strings the
// same value as received from the API will be provided.
type Conditions struct {
	TempC float32
	TempF float32

	FeelsLikeF float32
	FeelsLikeC float32

	PressureMb       float32
	RelativeHumidity string

	VisibilityKm float32

	Weather string

	Wind        string
	WindKph     float32
	WindGustKph float32
	WindDegrees float32
	WindDir     string

	ObservationTime time.Time
}

func (r Response) String() string {
	return fmt.Sprintf("{ CurrentConditions: %s }", r.CurrentConditions)
}
