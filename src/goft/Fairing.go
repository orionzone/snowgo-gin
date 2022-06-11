package goft

import "github.com/gin-gonic/gin"

//这是用来规范中间件代码和功能的接口
//Fairing整流罩。 用户保护卫星的

type Fairing interface {
	OnRequest(ctx *gin.Context) error
}
