package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/getNarutoList", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "ok",
			"data": "500",
		})
	})
	router.Run(":4000")
}
