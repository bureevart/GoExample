package dependency

import (
	"goexample/config"
	model "goexample/domain/models"
	contractRepository "goexample/domain/repository"
	database "goexample/infra/persistence/database"
	infraRepository "goexample/infra/persistence/repository"
)

func GetAreaInfoRepository(cfg *config.Config) contractRepository.AreaInfoRepository {
	//var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Tag"}}
	var preloads = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.AreaInfo](cfg, preloads)
}
