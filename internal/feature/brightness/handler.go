package brightness

import (
	"cmd/main.go/internal/model"
	"cmd/main.go/internal/repository/ylight"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type errorHandler interface {
	Handle(ctx *gin.Context, status int, err error)
}

type bulbGetter interface {
	Get(ip string) (model.Bulb, error)
}

type brightnessSetter interface {
	SetBrightness(loc string, brightness, duration int) (ylight.Response, error)
}

type Handler struct {
	errorHandler     errorHandler
	bulbGetter       bulbGetter
	brightnessSetter brightnessSetter
}

func NewHandler(
	errorHandler errorHandler,
	bulbGetter bulbGetter,
	brightnessSetter brightnessSetter,
) *Handler {
	return &Handler{
		errorHandler:     errorHandler,
		bulbGetter:       bulbGetter,
		brightnessSetter: brightnessSetter,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	var brightnessValue int
	err := json.Unmarshal(
		[]byte(ctx.PostForm("brightness")), &brightnessValue,
	)

	_, err = h.brightnessSetter.SetBrightness(ctx.PostForm("location"), brightnessValue, 1)
	if err != nil {
		h.errorHandler.Handle(ctx, http.StatusInternalServerError,
			fmt.Errorf("failed to toggle light: %w", err))

		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		h.errorHandler.Handle(ctx, http.StatusInternalServerError, err)

		return
	}

	err = tmpl.Execute(ctx.Writer, nil)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html", nil)

		return
	}
}
