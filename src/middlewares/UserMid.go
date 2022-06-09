package middlewares

import "fmt"

// 代表用户中间价
type UserMid struct {
}

func NewUserMid() *UserMid {
	return &UserMid{}
}

func (this *UserMid) OnRequest() error {
	fmt.Println("这是新的用户中间件")
	return fmt.Errorf("强迫写的错误")
}
