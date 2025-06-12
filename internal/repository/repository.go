package repository

import (
	"cmd/main.go/internal/model"
)

type Bulb interface {
	Get(id string) (model.Bulb, error)
	GetOfflineBulbs(onlineIDs []string) ([]model.Bulb, error)
}
