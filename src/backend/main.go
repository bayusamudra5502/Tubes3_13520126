package main

import (
	_ "github.com/bayusamudra5502/Tubes3_13520126/src/backend/db"
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