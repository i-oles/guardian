package main

import (
	"cmd/main.go/internal/repository/sqlite"
	"database/sql"
	"fmt"
	"log/slog"
)

func main() {
	db, err := sql.Open("sqlite3", "guardian.db")
	if err != nil {
		slog.Error(err.Error())
	}

	bulbRepo := sqlite.NewBulbsRepo(db, "bulbs")
	bulb, err := bulbRepo.Get("192.168.0.15")
	if err != nil {
		slog.Error("Error getting bulb: " + err.Error())
	}

	fmt.Printf("Bulb: %+v\n", bulb)
}
