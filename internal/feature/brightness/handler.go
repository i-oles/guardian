package brightness

import (
	"cmd/main.go/internal/model"
	"cmd/main.go/internal/repository/ylight"
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
	Set(ip string, brightness int) error
}

type Handler struct {
	errorHandler    errorHandler
	bulbGetter      bulbGetter
	brightnessSetter brightnessSetter
}

func NewHandler(
	errorHandler errorHandler,
	bulbGetter bulbGetter,
	brightnessSetter brightnessSetter,
) *Handler {
	return &Handler{
		errorHandler:    errorHandler,
		bulbGetter:      bulbGetter,
		brightnessSetter: brightnessSetter,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	loc := ctx.PostForm("location")
	bulb := ylight.YLight{Location: loc}

	_, err := bulb.Toggle()
	if err != nil {
		h.errorHandler.Handle(ctx, http.StatusInternalServerError,
			fmt.Errorf("failed to toggle light: %w", err)
		)

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
