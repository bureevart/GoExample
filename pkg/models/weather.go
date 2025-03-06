package models

// WeatherData - структура данных о погоде
// @Description WeatherData represents the weather information for a city
type WeatherData struct {
	ID uint `json:"id" gorm:"primaryKey"`
	// @Description The name of the city
	// @Example "London"
	City string `json:"city"`

	// @Description The current temperature in Celsius
	// @Example 25
	Temperature float32 `json:"temperature"`

	// @Description Weather condition (e.g. sunny, rainy)
	// @Example "Sunny"
	Condition string `json:"condition"`

	// @Description Timestamp of the weather data
	// @Example "2025-03-03T12:00:00Z"
	Timestamp string `json:"timestamp"`

	History []WeatherHistory `json:"history" gorm:"foreignKey:WeatherID"`
}

// WeatherHistory - история погоды для города
// @Description WeatherHistory keeps track of past weather conditions for a city
type WeatherHistory struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	WeatherID uint    `json:"weather_id" gorm:"index"` // Внешний ключ на WeatherData
	Temp      float32 `json:"temp"`
	Condition string  `json:"condition"`
	Recorded  string  `json:"recorded"`
}
