package router

import (
	favoriteRouter "github.com/czhi-bin/mini-tiktok-backend/biz/router/favorite"
	feedRouter "github.com/czhi-bin/mini-tiktok-backend/biz/router/feed"
	publishRouter "github.com/czhi-bin/mini-tiktok-backend/biz/router/publish"
	userRouter "github.com/czhi-bin/mini-tiktok-backend/biz/router/user"

	"github.com/gin-gonic/gin"	
)

func RegisterServices(r *gin.Engine) {
	favoriteRouter.RegisterRoutes(r)
	
	feedRouter.RegisterRoutes(r)
	
	publishRouter.RegisterRoutes(r)

	userRouter.RegisterRoutes(r)
}