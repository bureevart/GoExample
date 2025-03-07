package main

import (
	"goexample/api"
	"goexample/config"
	database "goexample/infra/persistence/database"
	"goexample/pkg/logging"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//	@title			Weather API
//	@version		1.0
//	@description	Weather information service
//	@host			localhost:8080
//	@BasePath		/api
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	api.InitServer(cfg)
}
