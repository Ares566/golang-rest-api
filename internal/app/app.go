package app

import (
	"rest-api-endpoints/application"
	"rest-api-endpoints/infrastructure/database"
	"rest-api-endpoints/infrastructure/logger"
	"rest-api-endpoints/internal/config"

	"github.com/joho/godotenv"
)

// all key app data
const (
	AppConfig   = iota
	dbConfig    = iota
	AccountsApp = iota
	TagsApp     = iota
	NotesApp    = iota
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
		
		tagRepository   = database.NewTagRepository(db.Conn)
		tagsApp         = application.NewTagApplication(tagRepository)
		notesRepository = database.NewNotesRepository(db.Conn)
		notesApp        = application.NewNotesApplication(notesRepository)
	)

	return map[int]interface{}{
		AppConfig:   appConfig,
		TagsApp:     tagsApp,
		NotesApp:    notesApp,
	}
}
