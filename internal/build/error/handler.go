package error

import (
	"cmd/main.go/internal/httpapi"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	stage     string
	isDebugOn bool
}

func New(stage string, isDebugOn bool) *Handler {
	return &Handler{
		stage:     stage,
		isDebugOn: isDebugOn,
	}
}

func (h *Handler) Handle(ctx *gin.Context, status int, err error) {
	if h.isDebugOn {
		slog.Error("API:", h.stage, status, "err", err)
	}

	ctx.JSON(status, httpapi.NewErrorBaseResponseBody(err))
}
