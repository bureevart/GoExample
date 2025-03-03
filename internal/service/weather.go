package service

import (
	"goexample/internal/repository"
)

// WeatherService обрабатывает логику работы с погодой.
type WeatherService struct {
	repo *repository.WeatherRepository
}

// NewWeatherService создает новый экземпляр WeatherService.
func NewWeatherService(repo *repository.WeatherRepository) *WeatherService {
	return &WeatherService{repo: repo}
}

// GetAllWeather получает все данные о погоде.
func (s *WeatherService) GetAllWeather() ([]repository.Weather, error) {
	return s.repo.GetAllWeather()
}

// CreateWeather создает новый объект погоды.
func (s *WeatherService) CreateWeather(weather *repository.Weather) error {
	return s.repo.CreateWeather(weather)
}
