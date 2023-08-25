package main

import (
	"github.com/gin-gonic/gin"

	"github.com/czhi-bin/mini-tiktok-backend/biz/dal/db"
	"github.com/czhi-bin/mini-tiktok-backend/biz/middleware/jwt"
	"github.com/czhi-bin/mini-tiktok-backend/biz/router"
)

func main() {
	db.Init()
	jwt.Init()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.RegisterServices(r)
	r.Run("0.0.0.0:18000")
}