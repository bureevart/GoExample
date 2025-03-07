package migration

import (
	"goexample/config"

	"goexample/pkg/logging"

	models "goexample/domain/models"

	database "goexample/infra/persistence/database"

	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up1() {
	database := database.GetDb()

	createTables(database)
}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	// Basic
	tables = addNewTable(database, models.AreaInfo{}, tables)
	tables = addNewTable(database, models.Device{}, tables)
	tables = addNewTable(database, models.GroupQuery{}, tables)
	tables = addNewTable(database, models.Node{}, tables)
	// Property
	tables = addNewTable(database, models.OrderBytes{}, tables)
	tables = addNewTable(database, models.Tag{}, tables)

	// User
	tables = addNewTable(database, models.TagsInGroupQuery{}, tables)
	tables = addNewTable(database, models.WeatherData{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func Down1() {
	// nothing
}
