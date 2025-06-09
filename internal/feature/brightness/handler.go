package brightness

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"cmd/main.go/internal/repository/ylight"
	"github.com/gin-gonic/gin"
)

type errorHandler interface {
	Handle(ctx *gin.Context, status int, err error)
}

type brightnessSetter interface {
	SetBrightness(loc string, brightness, duration int) (ylight.Response, error)
}

type Handler struct {
	errorHandler     errorHandler
	brightnessSetter brightnessSetter
}

func NewHandler(
	errorHandler errorHandler,
	brightnessSetter brightnessSetter,
) *Handler {
	return &Handler{
		errorHandler:     errorHandler,
		brightnessSetter: brightnessSetter,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	id := ctx.PostForm("id")
	location := ctx.PostForm("location")
	brightness := ctx.PostForm("brightness")

	if id == "" || location == "" || brightness == "" {
		h.errorHandler.Handle(ctx, http.StatusBadRequest,
			fmt.Errorf("missing required parameters"))
		return
	}

	slog.Info("Brightness request",
		"all_params", ctx.Request.PostForm,
		"id", id,
		"location", location,
		"brightness", brightness) // Dodaj to logowanie

	brightnessValue, err := strconv.Atoi(brightness)

	_, err = h.brightnessSetter.SetBrightness(location, brightnessValue, 1)
	if err != nil {
		h.errorHandler.Handle(ctx, http.StatusBadRequest,
			fmt.Errorf("invalid brightness value: %w", err))

		return
	}

	ctx.HTML(http.StatusOK, "brightness.tmpl", gin.H{
		"ID":         id,
		"Location":   location,
		"Brightness": brightnessValue,
	})
}
