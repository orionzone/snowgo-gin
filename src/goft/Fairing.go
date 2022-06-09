package goft

//这是用来规范中间件代码和功能的接口
//Fairing整流罩。 用户保护卫星的

type Fairing interface {
	OnRequest() error
}
