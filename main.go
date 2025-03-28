package main

import (
	"bakusoku/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/getNarutoList", controller.Getting)
	router.Run(":4000")
}
