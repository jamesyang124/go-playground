package basic

import "time"

// Bootcamp is exported
type Bootcamp struct {
	// Latitude of the event
	Lat float64
	// Longitude of the event
	Lon float64
	// Date of the event
	Date time.Time
}

var (
	// BootcampPtr exported
	BootcampPtr = &Bootcamp{Lat: 5.0, Lon: 6.0, Date: time.Now()}
)
