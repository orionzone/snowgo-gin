package main

import (
	"snowgo-gin/src/classes"
	"snowgo-gin/src/goft"
	. "snowgo-gin/src/middlewares"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	//classes.NewIndexClass(r).Build()
	//classes.NewUserClass(r).Build()
	//r.Run(":8080")
	goft.NewGoft().
		Attach(NewUserMid()).
		Mount("v1", classes.NewIndexClass()).
		Mount("v1", classes.NewUserClass()).
		Launch()
}
