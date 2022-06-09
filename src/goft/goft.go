package goft

import "github.com/gin-gonic/gin"

type Goft struct {
	*gin.Engine
	g *gin.RouterGroup
}

func NewGoft() *Goft {
	return &Goft{Engine: gin.New()}
}
func (this *Goft) Launch() {
	this.Run(":8080")
}
func (this *Goft) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Goft {
	this.g.Handle(httpMethod, relativePath, handlers...)
	return this
}
func (this *Goft) Mount(group string, classes ...IClass) *Goft {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
	}
	return this
}
