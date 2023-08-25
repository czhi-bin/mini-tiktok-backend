package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/czhi-bin/mini-tiktok-backend/biz/middleware/jwt"
	userModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/user"
	userService "github.com/czhi-bin/mini-tiktok-backend/biz/service/user"
)

// @router /douyin/user/register/ [POST]
func Register(c *gin.Context) {
	var err error
	var req userModel.UserRegisterRequest
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, userModel.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "Invalid parameters",
		})
		return
	}

	_, err = userService.NewService(c).Register(&req)
	if err != nil {
		c.JSON(http.StatusOK, userModel.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	jwt.JWTMiddleWare.LoginHandler(c)

	token := c.GetString("token")
	id, _ := c.Get("user_id")
	userId := id.(int64)

	c.JSON(http.StatusOK, userModel.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "Successfully registered",
		UserId: userId,
		Token:  token,
	})
}

// @router /douyin/user/login/ [POST]
func Login(c *gin.Context) {
	v, _ := c.Get("user_id")
	userId := v.(int64)
	token := c.GetString("token")
	c.JSON(http.StatusOK, userModel.UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "Successfully logged in",
		UserId: userId,
		Token:  token,
	})
}

// @router /douyin/user/ [GET]
func User(c *gin.Context) {
	var err error
	var req userModel.UserRequest
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, userModel.UserResponse{
			StatusCode: -1,
			StatusMsg:  "Invalid parameters",
		})
		return
	}

	userInfo, err := userService.NewService(c).GetUserInfo(&req)
	if err != nil {
		c.JSON(http.StatusOK, userModel.UserResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userModel.UserResponse{
		StatusCode: 0,
		StatusMsg:  "User retrieved successfully",
		User: userInfo,
	})
}
