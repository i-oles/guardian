package api

import "github.com/gin-gonic/gin"

type Responder interface {
	Success(ctx *gin.Context, data interface{})
	Error(ctx *gin.Context, errCode int, err error)
}
