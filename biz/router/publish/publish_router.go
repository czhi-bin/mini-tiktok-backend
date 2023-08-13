package publish

import (
	"github.com/gin-gonic/gin"

	publishHandler "github.com/czhi-bin/mini-tiktok-backend/biz/handler/publish"
)

// RegisterRoutes registers all routes of user module.
func RegisterRoutes(r *gin.Engine) {
	publishGroup := r.Group("/douyin/publish")

	publishGroup.POST("/action/", publishHandler.Action)
	publishGroup.GET("/list/", publishHandler.List)
}