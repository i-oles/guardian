package repository

import (
	"cmd/main.go/internal/model"
)

type BulbGetter interface {
	Get(id string) (model.Bulb, error)
}

type OfflineBulbsGetter interface {
	GetOfflineBulbs(onlineIDs []string) ([]model.Bulb, error)
}
