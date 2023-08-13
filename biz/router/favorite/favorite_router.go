package favorite

import (
	"github.com/gin-gonic/gin"

	favoriteHandler "github.com/czhi-bin/mini-tiktok-backend/biz/handler/favorite"
)

// RegisterRoutes registers all routes of user module.
func RegisterRoutes(r *gin.Engine) {
	favoriteGroup := r.Group("/douyin/favorite")

	favoriteGroup.POST("/action/", favoriteHandler.Action)
	favoriteGroup.GET("/list/", favoriteHandler.List)
}
