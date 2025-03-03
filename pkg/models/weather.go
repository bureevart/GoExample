package models

import "time"

// WeatherData - структура данных о погоде
// @Description WeatherData represents the weather information for a city
type WeatherData struct {
	// @Description The name of the city
	// @Example "London"
	City string `json:"city"`

	// @Description The current temperature in Celsius
	// @Example 25
	Temperature int `json:"temperature"`

	// @Description Weather condition (e.g. sunny, rainy)
	// @Example "Sunny"
	Condition string `json:"condition"`

	// @Description Timestamp of the weather data
	// @Example "2025-03-03T12:00:00Z"
	Timestamp time.Time `json:"timestamp"`
}
