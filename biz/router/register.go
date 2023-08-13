package router

import (
	userRouter "github.com/czhi-bin/mini-tiktok-backend/biz/router/user"

	"github.com/gin-gonic/gin"	
)

func RegisterServices(r *gin.Engine) {
	userRouter.RegisterRoutes(r)
}