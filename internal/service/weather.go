package service

import (
	"goexample/domain/models"
	"goexample/domain/repository"
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
func (s *WeatherService) GetAllWeather() ([]models.WeatherData, error) {
	return s.repo.GetAllWeather()
}

// CreateWeather – создание новой записи о погоде и добавление в историю
func (s *WeatherService) CreateWeather(data *models.WeatherData) error {
	return s.repo.CreateWeather(data)
}
