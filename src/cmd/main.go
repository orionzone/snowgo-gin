package main

import (
	"snowgo-gin/src/classes"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	classes.NewIndexClass(r).Build()
	r.Run(":8080")
}
