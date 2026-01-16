// Package weather provides tools 
// to forecast the current weather condition of various cities in Goblinocus.
package weather

var (
	// CurrentCondition represents the weather condition under the provided location.
	CurrentCondition string
	// CurrentLocation represents the city we are talking about now.
	CurrentLocation  string
)

// Forecast returns a string to represent the weather condition in current city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
