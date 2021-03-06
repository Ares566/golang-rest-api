package main

import (
	"fmt"
	"rest-api-endpoints/application"
	"rest-api-endpoints/infrastructure/logger"
	"rest-api-endpoints/interfaces"
	"rest-api-endpoints/internal/app"
	"rest-api-endpoints/internal/config"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "rest-api-endpoints/docs"
)

// @title Swagger For REST API
// @version 0.1
// @description Описание API

// @host localhost:8080
// @BasePath /
func main() {
	/* load app data and typed and define variable hanlder and inject data to constructor */
	var (
		appData   = app.New()
		appConfig = appData[app.AppConfig].(*config.AppConfig)

		// подключаем объявленные приложения и пробрасываем стратегию в interface

		tagsApp    = appData[app.TagsApp].(*application.TagApplication)
		tagHandler = interfaces.NewTagsREST(tagsApp)

		notesApp     = appData[app.NotesApp].(*application.NotesApplication)
		notesHandler = interfaces.NewNotesREST(notesApp)
	)

	/* define engine instance from gin framework */
	engine := gin.Default()
	if appConfig.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	CORSMiddleware(engine)

	

	if tagHandler != nil {
		engine.GET("/alltags", tagHandler.Fetch)
		engine.GET("/tagsbyproducts", tagHandler.FetchByProducts)
		engine.PUT("/tag", tagHandler.Update)
		engine.POST("/tag", tagHandler.Create)
		engine.DELETE("/tag", tagHandler.Delete)
	}

	if notesHandler != nil {
		engine.GET("/notes", notesHandler.Fetch)
		engine.PUT("/note", notesHandler.Update)
		engine.POST("/note", notesHandler.Create)
		engine.DELETE("/note", notesHandler.Delete)
	}

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	/* running server using engine instance from gin framework */
	if err := engine.Run(fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port)); err != nil {
		logger.Error(err)
	}
}

func CORSMiddleware(e *gin.Engine) {
	e.Use(gzip.Gzip(gzip.DefaultCompression))

	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           2 * time.Hour,
	}))

	e.Use(func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("Strict-Transport-Security", "true; includeSubDomains")
		c.Writer.Header().Set("X-Download-Options", "noopen")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Referrer-Policy", "no-referrer")
	})
}
