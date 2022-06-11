package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 代表用户中间价
type UserMid struct {
}

func NewUserMid() *UserMid {
	return &UserMid{}
}

//用gin.Context来获取一些请求参数
func (this *UserMid) OnRequest(ctx *gin.Context) error {
	fmt.Println("这是新的用户中间件")
	fmt.Println(ctx.Query("name"))
	//return fmt.Errorf("强迫写的错误")
	return nil

}
