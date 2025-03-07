package router

import (
	"goexample/api/handler"

	"goexample/config"

	"github.com/gin-gonic/gin"
)

const GetByFilterExp string = "/get-by-filter"

func AreaInfo(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewAreaInfoHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.GET("/", h.GetAll)
}
