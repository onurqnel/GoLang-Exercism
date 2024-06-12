/*
 * In Go, comments document code for `godoc`. They should start with the item's name and be placed before the item.
 */

// Package weather provides functionality to set and retrieve the current location and weather condition.
package weather

// CurrentCondition holds the weather condition for the current location.
var CurrentCondition string

// CurrentLocation holds the name of the current city.
var CurrentLocation string

// Forecast() sets the current location and weather condition, and returns a formatted string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
