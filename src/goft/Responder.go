package goft

import "github.com/gin-gonic/gin"

type Responder interface {
	RespondTo() gin.HandlerFunc
}
type StringResponder func(*gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, this(context))
	}
}
