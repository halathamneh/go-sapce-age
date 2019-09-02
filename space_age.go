package space

type Planet string

const (
	earthYearInSeconds = 31557600
)

var planetsRelativeYears = map[Planet]float32{
	Planet("Earth"):   1,
	Planet("Mercury"): 0.2408467,
	Planet("Venus"):   0.61519726,
	Planet("Mars"):    1.8808158,
	Planet("Jupiter"): 11.862615,
	Planet("Saturn"):  29.447498,
	Planet("Uranus"):  84.016846,
	Planet("Neptune"): 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	yearsInEarth := seconds / earthYearInSeconds
	return yearsInEarth / float64(planetsRelativeYears[planet])
}
