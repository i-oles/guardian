package home

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"cmd/main.go/internal/api"
	"cmd/main.go/internal/model"
	"cmd/main.go/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/julienrbrt/yeego/light/yeelight"
)

type Handler struct {
	bulbGetter         repository.BulbGetter
	offlineBulbsGetter repository.OfflineBulbsGetter
	apiResponder       api.Responder
}

func NewHandler(
	bulbGetter repository.BulbGetter,
	offlineBulbsGetter repository.OfflineBulbsGetter,
	apiResponder api.Responder,
) *Handler {
	return &Handler{
		bulbGetter:         bulbGetter,
		offlineBulbsGetter: offlineBulbsGetter,
		apiResponder:       apiResponder,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	onlineBulbs, err := yeelight.Discover(1 * time.Second)
	if err != nil {
		h.apiResponder.Error(ctx, http.StatusInternalServerError, err)

		return
	}

	onlineIDs := make([]string, len(onlineBulbs))
	for i, bulb := range onlineBulbs {
		onlineIDs[i] = bulb.ID
	}

	offlineBulbs, err := h.offlineBulbsGetter.GetOfflineBulbs(onlineIDs)
	if err != nil {
		h.apiResponder.Error(ctx, http.StatusInternalServerError, err)

		return
	}

	bulbStates, err := h.getBulbStates(onlineBulbs, offlineBulbs)
	if err != nil {
		h.apiResponder.Error(ctx, http.StatusInternalServerError, err)

		return
	}

	bulbStatesMapping := map[string][]model.BulbState{
		"Bulbs": bulbStates,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		h.apiResponder.Error(ctx, http.StatusInternalServerError, err)

		return
	}

	err = tmpl.Execute(ctx.Writer, bulbStatesMapping)
	if err != nil {
		h.apiResponder.Error(ctx, http.StatusInternalServerError, err)

		return
	}
}

func (h *Handler) getBulbStates(onlineBulbs []yeelight.Yeelight, offlineBulbs []model.Bulb) ([]model.BulbState, error) {
	var bulbStates []model.BulbState
	for _, bulb := range onlineBulbs {
		b, err := h.bulbGetter.Get(bulb.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get bulb with ID %s: %w", bulb.ID, err)
		}

		bulbState := model.BulbState{
			ID:         b.ID,
			Name:       b.Name,
			Location:   bulb.Location,
			State:      model.State(bulb.Power),
			Brightness: bulb.Bright,
		}

		bulbStates = append(bulbStates, bulbState)
	}

	for _, b := range offlineBulbs {
		bulbState := model.BulbState{
			ID:    b.ID,
			Name:  b.Name,
			State: model.Offline,
		}

		bulbStates = append(bulbStates, bulbState)
	}

	return bulbStates, nil
}
