package home

import (
	"cmd/main.go/internal/model"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"

	"github.com/julienrbrt/yeego/light/yeelight"
)

type bulbGetter interface {
	Get(id string) (model.Bulb, error)
	GetOfflineBulbs(onlineIDs []string) ([]model.Bulb, error)
}

type Handler struct {
	bulbGetter bulbGetter
}

func NewHandler(bulbGetter bulbGetter) *Handler {
	return &Handler{
		bulbGetter: bulbGetter,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	onlineBulbs, err := yeelight.Discover(1 * time.Second)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"error": err.Error(),
		})

		return
	}

	onlineIDs := make([]string, len(onlineBulbs))
	for i, bulb := range onlineBulbs {
		onlineIDs[i] = bulb.ID
	}

	offlineBulbs, err := h.bulbGetter.GetOfflineBulbs(onlineIDs)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"error": err.Error(),
		})

		return
	}

	var bulbStates []model.BulbState
	for _, bulb := range onlineBulbs {
		b, err := h.bulbGetter.Get(bulb.ID)
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"error": err.Error(),
			})

			return
		}

		isOn := bulb.Power == "on"

		bulbState := model.BulbState{
			ID:       b.ID,
			Name:     b.Name,
			Location: bulb.Location,
			IsOn:     &isOn,
		}

		bulbStates = append(bulbStates, bulbState)
	}

	for _, b := range offlineBulbs {
		bulbState := model.BulbState{
			ID:   b.ID,
			Name: b.Name,
		}

		bulbStates = append(bulbStates, bulbState)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html", nil)

		return
	}

	bb := map[string][]model.BulbState{
		"Bulbs": bulbStates,
	}

	err = tmpl.Execute(ctx.Writer, bb)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html", nil)

		return
	}
}
