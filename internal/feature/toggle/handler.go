package toggle

import (
	"fmt"
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

	resp, err := h.bulbToggler.Toggle(location)
	if err != nil {
		h.errorHandler.Handle(ctx, http.StatusInternalServerError,
			fmt.Errorf("failed to toggle light: %w", err))

		return
	}

	ctx.HTML(http.StatusOK, "toggle.tmpl", gin.H{
		"Name":     name,
		"Location": location,
		"State":    resp.Params.Power,
	})
}
