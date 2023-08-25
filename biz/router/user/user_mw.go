package user

import (
	"github.com/czhi-bin/mini-tiktok-backend/biz/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func LoginMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		jwt.JWTMiddleWare.LoginHandler,
	}
}

func UserMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		jwt.JWTMiddleWare.MiddlewareFunc(),
	}
}