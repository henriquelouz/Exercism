package space

// Planet as string
type Planet = string

// Seconds in an Earth year
const year = 31557600

// Map of planets and their orbital periods in Earth years
var orbitalPeriods = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns the given age in the given planet
func Age(seconds float64, planet Planet) float64 {
	secondsOnPlanet := year * orbitalPeriods[planet]
	return seconds / secondsOnPlanet
}
