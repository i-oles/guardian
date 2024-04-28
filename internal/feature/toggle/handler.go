package toggle

import (
	"cmd/main.go/internal/repository/ylight"
	"fmt"
	"net/http"

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
	isOn := ctx.PostForm("isOn")
	name := ctx.PostForm("name")

	_, err := h.bulbToggler.Toggle(location)
	if err != nil {
		h.errorHandler.Handle(ctx, http.StatusInternalServerError,
			fmt.Errorf("failed to toggle light: %w", err))

		return
	}

	var bulbIsOn bool

	if isOn == "true" {
		bulbIsOn = false
	} else {
		bulbIsOn = true
	}

	ctx.HTML(http.StatusOK, "button.tmpl", gin.H{
		"Name":     name,
		"Location": location,
		"IsOn":     bulbIsOn,
	})
}
