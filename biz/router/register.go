package router

import (
	feedRouter "github.com/czhi-bin/mini-tiktok-backend/biz/router/feed"
	userRouter "github.com/czhi-bin/mini-tiktok-backend/biz/router/user"

	"github.com/gin-gonic/gin"	
)

func RegisterServices(r *gin.Engine) {
	feedRouter.RegisterRoutes(r)
	
	userRouter.RegisterRoutes(r)
}