package toggle

import (
	"cmd/main.go/internal/httpapi"
	"cmd/main.go/internal/model"
	"fmt"
	"github.com/akominch/yeelight"
	"github.com/gin-gonic/gin"
	"net/http"
)

type toggleRequest struct {
	IP string `json:"ip"`
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

type Handler struct {
	errorHandler errorHandler
	bulbGetter   bulbGetter
}

func NewHandler(
	errorHandler errorHandler,
	bulbGetter bulbGetter,
) *Handler {
	return &Handler{
		errorHandler: errorHandler,
		bulbGetter:   bulbGetter,
	}
}

func (h *Handler) Handle(ctx *gin.Context) {
	var req toggleRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		h.errorHandler.Handle(ctx, http.StatusBadRequest,
			fmt.Errorf("could not bind toggleRequest: %w", err),
		)

		return
	}

	_, err := h.bulbGetter.Get(req.IP)
	if err != nil {
		h.errorHandler.Handle(ctx, http.StatusBadRequest,
			fmt.Errorf("could not find bulb %s", req.IP),
		)

		return
	}

	if err := toggleBulb(req.IP); err != nil {
		h.errorHandler.Handle(ctx, http.StatusInternalServerError,
			fmt.Errorf("could not toggle bulb %s", req.IP),
		)

		return
	}

	ctx.JSON(
		http.StatusOK, toggleResponse{
			Body: httpapi.NewOkBaseResponseBody(),
		},
	)
}

func toggleBulb(ip string) error {
	config := yeelight.BulbConfig{
		Ip:     ip,
		Effect: yeelight.Smooth,
	}

	bulb := yeelight.New(config)
	isOn, err := bulb.IsOn()
	if err != nil {
		return fmt.Errorf("could not check if bulb is on: %w", err)
	}

	if isOn {
		_, err := bulb.TurnOff()
		if err != nil {
			return fmt.Errorf("could not turn off bulb: %s, err: %w", ip, err)
		}
	} else {
		_, err := bulb.TurnOn()
		if err != nil {
			return fmt.Errorf("could not turn on bulb: %s, err: %w", ip, err)
		}
	}

	return nil
}
