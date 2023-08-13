package feed

import (
	"github.com/gin-gonic/gin"

	feedHandler "github.com/czhi-bin/mini-tiktok-backend/biz/handler/feed"
)

// RegisterRoutes registers all routes of user module.
func RegisterRoutes(r *gin.Engine) {
	feedGroup := r.Group("/douyin/feed")
	
	feedGroup.GET("", feedHandler.Feed)
}