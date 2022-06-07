package goft

import "github.com/gin-gonic/gin"

type Goft struct {
	*gin.Engine
}

func NewGoft() *Goft {
	return &Goft{Engine: gin.New()}
}
func (this *Goft) Launch() {
	this.Run(":8080")
}
func (this *Goft) Mount(classes ...IClass) *Goft {
	for _, class := range classes {
		class.Build(this)
	}
	return this
}
