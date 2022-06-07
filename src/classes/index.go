package classes

import (
	"snowgo-gin/src/goft"

	"github.com/gin-gonic/gin"
)

type IndexClass struct {
}

// NewIndexClass  is so-called constructor
func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

//This is our business method. The function name is optional
func (this *IndexClass) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "index ok",
		})
	}
}

// Build shadow this so that main func is clear
func (this *IndexClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", this.GetIndex())
}
