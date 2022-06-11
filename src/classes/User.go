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
func (this *UserClass) UserList(ctx *gin.Context) string {

	return "abc"
}
func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user", this.UserList)
}
