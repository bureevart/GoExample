package repository

import (
	"github.com/jinzhu/gorm"
)

// Weather представляет модель для таблицы weather в базе данных.
type Weather struct {
	ID          uint   `gorm:"primary_key"`
	City        string `gorm:"not null"`
	Temperature int    `gorm:"not null"`
	Condition   string `gorm:"not null"`
	Timestamp   string `gorm:"not null"`
}

// WeatherRepository описывает интерфейс для работы с данными погоды в базе данных.
type WeatherRepository struct {
	DB *gorm.DB
}

// NewWeatherRepository создает новый экземпляр репозитория.
func NewWeatherRepository(db *gorm.DB) *WeatherRepository {
	return &WeatherRepository{DB: db}
}

// GetAllWeather получает все данные о погоде из базы.
func (repo *WeatherRepository) GetAllWeather() ([]Weather, error) {
	var weather []Weather
	if err := repo.DB.Find(&weather).Error; err != nil {
		return nil, err
	}
	return weather, nil
}

// CreateWeather создает новый объект погоды в базе.
func (repo *WeatherRepository) CreateWeather(weather *Weather) error {
	return repo.DB.Create(weather).Error
}
