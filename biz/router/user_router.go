package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	userHandler "github.com/czhi-bin/mini-tiktok-backend/biz/handler/user"
)

// RegisterRoutes registers all routes of user module.
func RegisterRoutes(h *server.Hertz) {
	userGroup := h.Group("/douyin/user")
	
	userGroup.GET("/", userHandler.GetUserInfo)
	userGroup.POST("/register", userHandler.Register)
	userGroup.POST("/login", userHandler.Login)
}