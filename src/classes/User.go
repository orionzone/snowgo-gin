package classes

import (
	"snowgo-gin/src/goft"

	"github.com/gin-gonic/gin"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}
func (this *UserClass) UserList() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "success",
		})
	}
}
func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user", this.UserList())
}
