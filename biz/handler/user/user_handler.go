package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/czhi-bin/mini-tiktok-backend/biz/middleware/jwt"
	userModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/user"
	"github.com/czhi-bin/mini-tiktok-backend/biz/model/common"
	userService "github.com/czhi-bin/mini-tiktok-backend/biz/service/user"
)

// @router /douyin/user/register/ [POST]
func Register(c *gin.Context) {
	var err error
	var req userModel.UserRegisterRequest
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: -1,
			StatusMsg:  "Invalid parameters",
		})
		return
	}

	userId, err := userService.NewService(c).Register(&req)
	if err != nil {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	token, err := jwt.GenerateToken(userId)
	if err != nil {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userModel.UserRegisterResponse{
		CommonResponse: &common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "Successfully registered",
		},
		UserAuth: &common.UserAuth{
			UserId: userId,
			Token:  token,
		},
	})
}

// @router /douyin/user/login/ [POST]
func Login(c *gin.Context) {
	var err error
	var req userModel.UserLoginRequest
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: -1,
			StatusMsg:  "Invalid parameters",
		})
		return
	}

	userAuth, err := userService.NewService(c).Login(&req)
	if err != nil {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userModel.UserLoginResponse{
		CommonResponse: &common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "Successfully logged in",
		},
		UserAuth: userAuth,
	})
}

// @router /douyin/user/ [GET]
func GetUserInfo(c *gin.Context) {

}
