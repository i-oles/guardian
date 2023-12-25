package httpapi

import "time"

type BaseResponseBody struct {
	Ok        bool      `json:"ok"`
	ErrorMsg  string    `json:"error_msg,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

func NewOkBaseResponseBody() BaseResponseBody {
	return BaseResponseBody{
		Ok:        true,
		Timestamp: time.Now(),
	}
}

func NewErrorBaseResponseBody(err error) BaseResponseBody {
	return BaseResponseBody{
		Ok:        false,
		Timestamp: time.Now(),
		ErrorMsg:  err.Error(),
	}
}
