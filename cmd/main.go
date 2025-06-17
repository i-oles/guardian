package main

import (
	"database/sql"
	"log"
	"log/slog"

	"cmd/main.go/internal/api/http"
	"cmd/main.go/internal/bulb/controller"
	"cmd/main.go/internal/config"
	"cmd/main.go/internal/handler/brightness"
	"cmd/main.go/internal/handler/home"
	"cmd/main.go/internal/handler/toggle"
	"cmd/main.go/internal/repository/sqlite"
	configuration "cmd/main.go/pkg/config"
	"github.com/gin-gonic/gin"
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

	bulbRepo := sqlite.NewBulbRepo(db, cfg.BulbCollName)
	bulbController := controller.NewYeeLight()

	homeHandler := home.NewHandler(
		bulbRepo, bulbRepo,
		http.NewAPIResponder("/home", cfg.Logging),
	)

	GETs := map[string]gin.HandlerFunc{
		"/": homeHandler.Handle,
	}

	toggleHandler := toggle.NewHandler(
		http.NewAPIResponder("/toggle", cfg.Logging),
		bulbController)

	brightnessHandler := brightness.NewHandler(
		http.NewAPIResponder("/brightness", cfg.Logging),
		bulbController)

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
