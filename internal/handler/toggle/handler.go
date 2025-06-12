package toggle

import (
	"net/http"

	"cmd/main.go/internal/api"
	"cmd/main.go/internal/bulb"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	apiResponder   api.Responder
	bulbController bulb.Controller
}

func NewHandler(
	apiResponder api.Responder,
	bulbController bulb.Controller,
) *Handler {
	return &Handler{
		apiResponder:   apiResponder,
		bulbController: bulbController,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	location := ctx.PostForm("location")
	name := ctx.PostForm("name")
	id := ctx.PostForm("id")
	brightness := ctx.PostForm("brightness")

	resp, err := h.bulbController.Toggle(location)
	if err != nil {
		h.apiResponder.Error(ctx, http.StatusInternalServerError, err)

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
