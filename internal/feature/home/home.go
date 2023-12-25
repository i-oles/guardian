package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Hello, world!",
	})
}
