package toggle

import (
	"fmt"
	"log/slog"
	"net/http"

	"cmd/main.go/internal/repository/ylight"
	"github.com/gin-gonic/gin"
)

type errorHandler interface {
	Handle(ctx *gin.Context, status int, err error)
}

type bulbToggler interface {
	Toggle(loc string) (ylight.Response, error)
}

type Handler struct {
	errorHandler errorHandler
	bulbToggler  bulbToggler
}

func NewHandler(
	errorHandler errorHandler,
	bulbToggler bulbToggler,
) *Handler {
	return &Handler{
		errorHandler: errorHandler,
		bulbToggler:  bulbToggler,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	location := ctx.PostForm("location")
	name := ctx.PostForm("name")
	id := ctx.PostForm("id")
	brightness := ctx.PostForm("brightness")

	slog.Info("Bulb state",
		"id", id,
		"name", name,
		"location", location,
		"brightness", brightness)

	resp, err := h.bulbToggler.Toggle(location)
	if err != nil {
		h.errorHandler.Handle(ctx, http.StatusInternalServerError,
			fmt.Errorf("failed to toggle light: %w", err))

		return
	}

	ctx.HTML(http.StatusOK, "toggle.tmpl", gin.H{
		"ID":         id,
		"Name":       name,
		"Location":   location,
		"State":      resp.Params.Power,
		"Brightness": brightness,
	})
}
