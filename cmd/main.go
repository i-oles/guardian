package main

import (
	errorHandler "cmd/main.go/internal/build/error"
	"cmd/main.go/internal/config"
	"cmd/main.go/internal/feature/brightness"
	"cmd/main.go/internal/feature/home"
	"cmd/main.go/internal/feature/toggle"
	"cmd/main.go/internal/repository/sqlite"
	"cmd/main.go/internal/repository/ylight"
	configuration "cmd/main.go/pkg/config"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
)

func main() {
	var cfg config.Configuration

	err := configuration.GetConfig("./config", &cfg)
	if err != nil {
		slog.Error(err.Error())
	}

	log.Println(cfg.Pretty())

	db, err := sql.Open("sqlite3", cfg.DBFileName)
	if err != nil {
		slog.Error(err.Error())
	}

	bulbRepo := sqlite.NewBulbsRepo(db, cfg.BulbCollName)
	yeeLight := ylight.NewYLight()

	homeHandler := home.NewHandler(bulbRepo)

	GETs := map[string]gin.HandlerFunc{
		"/": homeHandler.Handle,
	}

	toggleHandler := toggle.NewHandler(
		errorHandler.New("toggle", cfg.IsDebugOn),
		yeeLight)

	brightnessHandler := brightness.NewHandler(
		errorHandler.New("brightness", cfg.IsDebugOn),
		bulbRepo,
		yeeLight)

	POSTs := map[string]gin.HandlerFunc{
		"/toggle":     toggleHandler.Handle,
		"/brightness": brightnessHandler.Handle,
	}

	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	webApp := router.Group("/")

	for route, handler := range GETs {
		webApp.GET(route, handler)
	}

	for route, handler := range POSTs {
		webApp.POST(route, handler)
	}

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
