// Package weather provides tools to display the current weather forecast for a city.
package weather

// CurrentCondition represents a weather condition as a string.
var CurrentCondition string

// CurrentLocation represents a city name as a string.
var CurrentLocation string

// Forecast takes a city name and weather condition as strings and returns a string describing the weather in that city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
