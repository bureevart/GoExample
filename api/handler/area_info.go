package handler

import (
	"goexample/api/dto"
	_ "goexample/api/helper"
	"goexample/config"
	"goexample/dependency"
	"goexample/usecase"

	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	usecase *usecase.AreaInfoUsecase
}

func NewAreaInfoHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{
		usecase: usecase.NewAreaInfoUsecase(cfg, dependency.GetAreaInfoRepository(cfg)),
	}
}

// CreateAreaInfo godoc
// @Summary Create a AreaInfo
// @Description Create a AreaInfo
// @Tags AreaInfo
// @Accept json
// @produces json
// @Param Request body dto.CreateAreaInfoRequest true "Create a AreaInfo"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.AreaInfoResponse} "AreaInfo response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/areainfo/ [post]
func (h *CityHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateAreaInfo, dto.ToAreaInfoResponse, h.usecase.Create)
}

// UpdateAreaInfo godoc
// @Summary Update a AreaInfo
// @Description Update a AreaInfo
// @Tags AreaInfo
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateAreaInfoRequest true "Update a AreaInfo"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CityResponse} "AreaInfo response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/areainfo/{id} [put]
func (h *CityHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateAreaInfo, dto.ToAreaInfoResponse, h.usecase.Update)
}

// DeleteAreaInfo godoc
// @Summary Delete a AreaInfo
// @Description Delete a AreaInfo
// @Tags AreaInfo
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/areainfo/{id} [delete]
func (h *CityHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetAreaInfo godoc
// @Summary Get a AreaInfo
// @Description Get a AreaInfo
// @Tags AreaInfo
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.AreaInfoResponse} "AreaInfo response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/areainfo/{id} [get]
func (h *CityHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToAreaInfoResponse, h.usecase.GetById)
}

// GetAllAreaInfo godoc
// @Summary Get a AreaInfos
// @Description Get a AreaInfos
// @Tags AreaInfos
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.AreaInfoResponse} "AreaInfos response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/areainfo/ [get]
func (h *CityHandler) GetAll(c *gin.Context) {
	GetById(c, dto.ToAreaInfoResponse, h.usecase.GetById)
}
