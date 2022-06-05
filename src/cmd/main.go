package main

import "github.com/gin-gonic/gin"

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Handle("GET", "/", func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "success"})
	})
	r.Run(":8080")
}
