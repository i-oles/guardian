package toggle

import (
	"cmd/main.go/internal/repository/ylight"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorHandler interface {
	Handle(ctx *gin.Context, status int, err error)
}

type bulbToggler interface {
	Toggle() (ylight.Response, error)
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
	bulb := ylight.YLight{Location: ctx.PostForm("location")}

	_, err := bulb.Toggle()
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
