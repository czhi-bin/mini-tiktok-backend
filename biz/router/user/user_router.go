package user

import (
	"github.com/gin-gonic/gin"

	userHandler "github.com/czhi-bin/mini-tiktok-backend/biz/handler/user"
)

// RegisterRoutes registers all routes of user module.
func RegisterRoutes(r *gin.Engine) {
	userGroup := r.Group("/douyin/user")

	userGroup.GET("/", append(UserMw(), userHandler.User)...)
	userGroup.POST("/register/", userHandler.Register)
	userGroup.POST("/login/", append(LoginMw(), userHandler.Login)...)
}
