package http

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type APIResponder struct {
	Ok        bool        `json:"ok"`
	ErrorMsg  string      `json:"error_msg,omitempty"`
	ErrorCode int         `json:"error_code,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
	Logging   bool        `json:"logging"`
	Endpoint  string      `json:"endpoint,omitempty"`
}

func NewAPIResponder(endpoint string, logging bool) APIResponder {
	return APIResponder{
		Endpoint: endpoint,
		Logging:  logging,
	}
}

// TODO: use it and check if it's needed
func (a APIResponder) Success(
	ctx *gin.Context,
	data interface{},
) {
	if a.Logging {
		slog.Info("API:",
			slog.Bool("Ok", a.Ok),
			slog.String("endpoint", a.Endpoint),
			slog.String("data", fmt.Sprintf("%+v", data)),
		)
	}

	resp := APIResponder{
		Ok:        true,
		Endpoint:  a.Endpoint,
		Data:      data,
		Timestamp: time.Now(),
	}

	ctx.JSON(http.StatusOK, resp)
}

func (a APIResponder) Error(ctx *gin.Context, errCode int, err error) {
	if a.Logging {
		slog.Error("API:",
			slog.Bool("Ok", false),
			slog.String("endpoint", a.Endpoint),
			slog.String("err", err.Error()),
		)
	}

	resp := APIResponder{
		Ok:        false,
		Timestamp: time.Now(),
		ErrorCode: errCode,
		ErrorMsg:  err.Error(),
	}

	ctx.AbortWithStatusJSON(errCode, resp)
}
