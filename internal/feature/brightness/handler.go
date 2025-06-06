package brightness

import (
	"cmd/main.go/internal/repository/ylight"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	brightnessValue, err := strconv.Atoi(brightness)

	_, err = h.brightnessSetter.SetBrightness(location, brightnessValue, 1)
	if err != nil {
		h.errorHandler.Handle(ctx, http.StatusInternalServerError,
			fmt.Errorf("failed to toggle light: %w", err))

		return
	}

	ctx.HTML(http.StatusOK, "brightness.tmpl", gin.H{
		"Location":   location,
		"Brightness": brightnessValue,
	})
}
