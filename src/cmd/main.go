package main

import (
	"snowgo-gin/src/classes"
	"snowgo-gin/src/goft"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	//classes.NewIndexClass(r).Build()
	//classes.NewUserClass(r).Build()
	//r.Run(":8080")
	goft.NewGoft().
		Mount("v1", classes.NewIndexClass()).
		Mount("v2", classes.NewUserClass()).
		Launch()
}
