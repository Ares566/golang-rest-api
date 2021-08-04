package app

import (
	"rest-api-endpoints/application"
	"rest-api-endpoints/infrastructure/database"
	"rest-api-endpoints/internal/config"

	"rest-api-endpoints/util/logger"

	"github.com/joho/godotenv"
)

// all key app data
const (
	AppConfig = iota
	dbConfig  = iota
	TagsApp  = iota
)

// New is
func New() map[int]interface{} {

	//load environment variable
	if err := godotenv.Load(); err != nil {
		logger.Error(err)
	}

	var (
		appConfig = config.NewAppConfig()
		dbConfig  = config.NewDatabaseConfig()
	)

	
	// connect to database
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		logger.Error(err)
	}


	// dependency injection
	// TODO может uber-go/dig ?
	var (
		tagRepository = database.NewTagRepository(db.Conn)
		tagsApp      = application.NewTagApplication(tagRepository)
	)

	
	return map[int]interface{}{
		AppConfig: appConfig,
		TagsApp: tagsApp,
	}
}
