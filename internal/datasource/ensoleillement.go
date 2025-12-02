package datasource

import "time"

var EnsoleillementHeure map[time.Month]float64 = map[time.Month]float64{
	time.January:   8,
	time.February:  7,
	time.March:     7,
	time.April:     6,
	time.May:       6,
	time.June:      6,
	time.July:      6,
	time.August:    6,
	time.September: 7,
	time.October:   7,
	time.November:  7,
	time.December:  8,
}
