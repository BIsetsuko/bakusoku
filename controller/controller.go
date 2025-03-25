package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getting(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"category": "木の葉隠れ",
		"name": "うずまき ナルト", // 仮のデータ
	})
}