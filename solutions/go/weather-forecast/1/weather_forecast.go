// Package weather provides weather forecasting functionalities.
package weather

var (
	// CurrentCondition - variable to hold current weather condition.
	CurrentCondition string
	// CurrentLocation -  variable to hold current location.
	CurrentLocation string
)

// Forecast returns a string combining the city and its current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
