package toggle

import (
	"fmt"
	"net/http"

	"cmd/main.go/internal/api"
	"cmd/main.go/internal/bulb"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	apiResponder api.Responder
	toggler      bulb.Toggler
}

func NewHandler(
	apiResponder api.Responder,
	toggler bulb.Toggler,
) *Handler {
	return &Handler{
		apiResponder: apiResponder,
		toggler:      toggler,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	location := ctx.PostForm("location")
	name := ctx.PostForm("name")
	id := ctx.PostForm("id")
	brightness := ctx.PostForm("brightness")

	if id == "" || location == "" || brightness == "" || name == "" {
		h.apiResponder.Error(ctx, http.StatusBadRequest,
			fmt.Errorf("missing required parameters"))

		return
	}

	resp, err := h.toggler.Toggle(location)
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
