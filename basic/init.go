package basic

import "time"

// Go supports the new expression to allocate a zeroed value of the requested type and to return a pointer to it.

var (
	// InitBootcamp ...
	// directly call same package public methods from other file
	// allocate a zeroed value of Bootcamp
	InitBootcamp             = new(Bootcamp)
	emptyBootcamp            = Bootcamp{}
	withDefaultValueBootcamp = Bootcamp{Lat: 1.0, Lon: 2.0, Date: time.Now()}
)
