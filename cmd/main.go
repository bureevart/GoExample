package main

import (
	"fmt"
	"log"
	"net/http"

	"goexample/internal/handler"
	"goexample/internal/repository"
	"goexample/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "goexample/cmd/docs" // Импортируем docs, чтобы swagger.json был доступен

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Weather API
// @version 1.0
// @description Weather information service
// @host localhost:8080
// @BasePath /api
func main() {
	// Строка подключения
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable" // Подключение к базе
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer db.Close()

	// Подключение и создание базы данных
	if err := createDatabaseIfNotExists(db); err != nil {
		log.Fatal("Unable to create database:", err)
	}

	// Подключаемся к созданной базе данных
	dsn = "host=localhost user=postgres password=postgres dbname=weatherdb port=5432 sslmode=disable"
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Unable to connect to weatherdb:", err)
	}
	defer db.Close()

	// Автоматическая миграция
	if err := db.AutoMigrate(&repository.Weather{}).Error; err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Migrations applied successfully!")

	weatherRepo := repository.NewWeatherRepository(db)
	weatherService := service.NewWeatherService(weatherRepo)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Роуты
	r.Route("/api", func(r chi.Router) {
		// Получение погоды
		// @Summary Get weather information
		// @Description Get all weather data
		// @Tags weather
		// @Accept  json
		// @Produce  json
		// @Success 200 {array} repository.Weather
		// @Router /api/weather [get]
		r.Get("/weather", weatherHandler.GetWeather)
		// Создание записи о погоде
		// @Summary Create weather information
		// @Description Create new weather data
		// @Tags weather
		// @Accept  json
		// @Produce  json
		// @Param weather body repository.Weather true "Weather data"
		// @Success 200 {object} repository.Weather
		// @Router /api/weather [post]
		r.Post("/weather", weatherHandler.CreateWeather)
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}

// createDatabaseIfNotExists проверяет наличие базы данных и создает ее, если она отсутствует
func createDatabaseIfNotExists(db *gorm.DB) error {
	var result []struct{} // Срез структур, поскольку нам не нужно получать данные, просто проверка наличия результата

	// Проверка существования базы данных
	if err := db.Raw("SELECT 1 FROM pg_database WHERE datname = 'weatherdb'").Scan(&result).Error; err != nil {
		return err
	}

	if len(result) == 0 {
		// База данных не существует, создаем
		if err := db.Exec("CREATE DATABASE weatherdb").Error; err != nil {
			return err
		}
		fmt.Println("Database weatherdb created successfully!")
	}

	return nil
}
