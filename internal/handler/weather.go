package handler

import (
	"encoding/json"
	"net/http"

	"goexample/internal/service"
	"goexample/pkg/models"
)

// WeatherHandler обрабатывает запросы, связанные с погодой.
type WeatherHandler struct {
	service *service.WeatherService
}

// NewWeatherHandler создает новый экземпляр WeatherHandler.
func NewWeatherHandler(service *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{service: service}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	weather, err := h.service.GetAllWeather()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(weather)
}

func (h *WeatherHandler) CreateWeather(w http.ResponseWriter, r *http.Request) {
	var weather models.WeatherData
	if err := json.NewDecoder(r.Body).Decode(&weather); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateWeather(&weather); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(weather)
}
