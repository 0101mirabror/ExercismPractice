

// Package weather package describe in current country weather forecast.
package weather


//CurrentCondition is current weather in location.
var CurrentCondition string
//CurrentLocation is the location of weather.
var CurrentLocation string


// Forecast function helps to forecast the weather for city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
