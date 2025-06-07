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
	location := ctx.PostForm("location")
	brightness := ctx.PostForm("brightness")
	slog.Info("location:", slog.String("location", location))
	slog.Info("brightness:", slog.String("brightness", brightness))

	brightnessValue, err := strconv.Atoi(brightness)

	_, err = h.brightnessSetter.SetBrightness(location, brightnessValue, 1)
	if err != nil {
		h.errorHandler.Handle(ctx, http.StatusBadRequest,
			fmt.Errorf("invalid brightness value: %w", err))

		return
	}

	ctx.HTML(http.StatusOK, "brightness.tmpl", gin.H{
		"Location":   location,
		"Brightness": brightnessValue,
	})
}
