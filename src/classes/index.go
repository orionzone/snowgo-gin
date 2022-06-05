package classes

import "github.com/gin-gonic/gin"

type IndexClass struct {
	*gin.Engine //this is what's gin new() get
}

// NewIndexClass  is so-called constructor
func NewIndexClass(e *gin.Engine) *IndexClass {
	return &IndexClass{Engine: e}
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
func (this *IndexClass) Build() {
	this.Handle("GET", "/", this.GetIndex())
}
