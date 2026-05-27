// Package weather provies tools to forecast weather.
package weather

var (
    //  CurrentCondition represents a certain weather pattern.
	CurrentCondition string
    //  CurrentLocation represents a location.
	CurrentLocation  string
)

//  Forecast returns a string value equal to the city and it's weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
