package brightness

import (
	"fmt"
	"net/http"
	"strconv"

	"cmd/main.go/internal/api"
	"cmd/main.go/internal/bulb"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	brightnessSetter bulb.BrightnessSetter
	apiResponder     api.Responder
}

func NewHandler(
	apiResponder api.Responder,
	brightnessSetter bulb.BrightnessSetter,
) *Handler {
	return &Handler{
		apiResponder:     apiResponder,
		brightnessSetter: brightnessSetter,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	id := ctx.PostForm("id")
	location := ctx.PostForm("location")
	brightness := ctx.PostForm("brightness")

	if id == "" || location == "" || brightness == "" {
		h.apiResponder.Error(ctx, http.StatusBadRequest,
			fmt.Errorf("missing required parameters"))

		return
	}

	brightnessValue, err := strconv.Atoi(brightness)

	_, err = h.brightnessSetter.SetBrightness(location, brightnessValue, 1)
	if err != nil {
		h.apiResponder.Error(ctx, http.StatusBadRequest,
			fmt.Errorf("invalid brightness value: %w", err))

		return
	}

	ctx.HTML(http.StatusOK, "brightness.tmpl", gin.H{
		"ID":         id,
		"Location":   location,
		"Brightness": brightnessValue,
	})
}
