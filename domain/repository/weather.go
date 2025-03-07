package repository

import (
	"goexample/domain/models"

	"github.com/jinzhu/gorm"
)

// WeatherRepository описывает интерфейс для работы с данными погоды в базе данных.
type WeatherRepository struct {
	DB *gorm.DB
}

// NewWeatherRepository создает новый экземпляр репозитория.
func NewWeatherRepository(db *gorm.DB) *WeatherRepository {
	return &WeatherRepository{DB: db}
}

// GetAllWeather получает все данные о погоде из базы.
func (repo *WeatherRepository) GetAllWeather() ([]models.WeatherData, error) {
	var weather []models.WeatherData
	if err := repo.DB.Find(&weather).Error; err != nil {
		return nil, err
	}
	return weather, nil
}

// CreateWeather создает новый объект погоды в базе.
func (repo *WeatherRepository) CreateWeather(weather *models.WeatherData) error {
	return repo.DB.Create(weather).Error
}

// AddWeatherHistory – добавление записи в историю погоды
func (repo *WeatherRepository) AddWeatherHistory(history *models.WeatherHistory) error {
	return repo.DB.Create(history).Error
}
