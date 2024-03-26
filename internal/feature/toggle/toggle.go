package toggle

import (
	"cmd/main.go/internal/repository/ylight"
	"fmt"
	"html/template"
	"net/http"

	"cmd/main.go/internal/httpapi"
	"cmd/main.go/internal/model"
	"github.com/gin-gonic/gin"
)

type toggleRequest struct {
	Location string `json:"location"`
}

type toggleResponse struct {
	Body httpapi.BaseResponseBody
}

type errorHandler interface {
	Handle(ctx *gin.Context, status int, err error)
}

type bulbGetter interface {
	Get(ip string) (model.Bulb, error)
}

type yeeLightToggler interface {
	Toggle() (ylight.Response, error)
}

type Handler struct {
	errorHandler    errorHandler
	bulbGetter      bulbGetter
	yeeLightToggler yeeLightToggler
}

func NewHandler(
	errorHandler errorHandler,
	bulbGetter bulbGetter,
	yeeLightToggler yeeLightToggler,
) *Handler {
	return &Handler{
		errorHandler:    errorHandler,
		bulbGetter:      bulbGetter,
		yeeLightToggler: yeeLightToggler,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	loc := ctx.PostForm("location")
	bulb := ylight.YLight{Location: loc}

	_, err := bulb.Toggle()
	if err != nil {
		h.errorHandler.Handle(ctx, http.StatusInternalServerError,
			fmt.Errorf("could not toggle bulb %s", bulb.Location),
		)

		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html", nil)

		return
	}

	err = tmpl.Execute(ctx.Writer, nil)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "index.html", nil)

		return
	}
}
