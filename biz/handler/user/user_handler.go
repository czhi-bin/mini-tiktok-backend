package user

import (
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	
	"github.com/czhi-bin/mini-tiktok-backend/biz/middleware/jwt"
	model "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/user"
	"github.com/czhi-bin/mini-tiktok-backend/biz/model/common"
	service "github.com/czhi-bin/mini-tiktok-backend/biz/service"
)

// @router /douyin/user/register/ [POST]
func Register(c *gin.Context) {
	var err error
	var req model.UserRegisterRequest
	err = c.BindQuery(&req)
	if err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: -1,
			StatusMsg: "Invalid parameters",
		})
		return
	}

	var userId int64
	userId, err = service.NewUserService(c).Register(&req)
	if err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: -1,
			StatusMsg: err.Error(),
		})
		return
	}

	var token string
	token, err = jwt.GenerateToken(userId)
	if err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: -1,
			StatusMsg: err.Error(),
		})
		return
	}

	fmt.Println("token IS ???", token)
	c.JSON(http.StatusOK, model.UserRegisterResponse{
		CommonResponse: &common.CommonResponse{
			StatusCode: 0,
			StatusMsg: 	"Successfully registered",
		},
		UserAuth: &common.UserAuth{
			UserId: 	userId, 		
			Token: 		token,
		},
	})
}

// @router /douyin/user/login/ [POST]
func Login(c *gin.Context) {

}

// @router /douyin/user/ [GET]
func GetUserInfo(c *gin.Context) {
	
}