package main

import (
	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
)

func main(){
	env.Load()
	
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
			"message": "success",
		})
	})

	r.Run()
}