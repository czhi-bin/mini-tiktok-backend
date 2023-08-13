package user

import (
	"net/http"
	"fmt"

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

	userId, err := userService.NewService(c).Register(&req)
	if err != nil {
		c.JSON(http.StatusOK, userModel.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	token, err := jwt.GenerateToken(userId)
	if err != nil {
		c.JSON(http.StatusOK, userModel.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userModel.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "Successfully registered",
		UserId: userId,
		Token:  token,
	})
}

// @router /douyin/user/login/ [POST]
func Login(c *gin.Context) {
	fmt.Println("in login handler")
	var err error
	var req userModel.UserLoginRequest
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, userModel.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "Invalid parameters",
		})
		return
	}

	fmt.Println("afer binding query")
	fmt.Println("req: ", req)
	fmt.Println("username: ", req.Username)
	fmt.Println("password: ", req.Password)
	userAuth, err := userService.NewService(c).Login(&req)
	if err != nil {
		c.JSON(http.StatusOK, userModel.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	fmt.Println("after login service")
	fmt.Println("userAuth: ", userAuth)
	fmt.Println("userId: ", userAuth.UserId)
	fmt.Println("token: ", userAuth.Token)
	// c.JSON(http.StatusOK, userModel.UserLoginResponse{
	// 	
	// 		StatusCode: 0,
	// 	},
	// 	UserAuth: userAuth,
	// })
	c.JSON(http.StatusOK, userModel.UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "Successfully logged in",
		UserId: userAuth.UserId,
		Token:  userAuth.Token,
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

	user, err := userService.NewService(c).GetUserInfo(&req)
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
		User: user,
	})
}
