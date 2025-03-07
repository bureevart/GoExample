package api

import (
	"fmt"

	// "goexample/docs"
	"goexample/pkg/logging"

	"goexample/config"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"goexample/api/middleware"
	"goexample/api/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

// var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()
	RegisterValidators()

	r.Use(middleware.DefaultStructuredLogger(cfg))
	r.Use(middleware.Cors(cfg))

	r.Use(gin.Logger(), gin.CustomRecovery(middleware.ErrorHandler) /*middleware.TestMiddleware()*/, middleware.LimitByRequest())

	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)
	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		// Base
		areainfos := v1.Group("/areainfos")

		// Base
		router.AreaInfo(areainfos, cfg)

		r.Static("/static", "./uploads")

		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
}

func RegisterValidators() {
	// val, ok := binding.Validator.Engine().(*validator.Validate)
	// if ok {
	// 	err = val.RegisterValidation("password", validation.PasswordValidator, true)
	// 	if err != nil {
	// 		logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
	// 	}
	// }
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
