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
func (this *Goft) Handle(httpMethod, relativePath string, handler interface{}) *Goft {
	if h := Convert(handler); h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	}
	return this
}

//代表加入中间件
func (this *Goft) Attach(f Fairing) *Goft {
	this.Use(func(context *gin.Context) {
		err := f.OnRequest(context)
		if err != nil {
			context.AbortWithStatusJSON(
				400, gin.H{"error": err.Error()})
		} else {
			context.Next()
		}

	})
	return this
}
func (this *Goft) Mount(group string, classes ...IClass) *Goft {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
	}
	return this
}
