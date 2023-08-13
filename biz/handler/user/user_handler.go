package user

import (
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	
	model "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/user"
	service "github.com/czhi-bin/mini-tiktok-backend/biz/service"
)

// @router /douyin/user/register/ [POST]
func Register(c *gin.Context) {
	var err error
	var req model.UserRegisterRequest
	err = c.BindQuery(&req)
	if err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusOK, model.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg: "Invalid parameters",
		})
		return
	}

	var user_id int64
	user_id, err = service.NewUserService(c).GinRegister(&req)
	if err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusOK, model.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg: err.Error(),
		})
		return
	}

	// TODO: replace with proper response
	c.JSON(http.StatusOK, model.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg: 	"Successfully registered",
		UserId: 	user_id, 		
		Token: 		"token",		// replace with proper token
	})
}

// @router /douyin/user/login/ [POST]
func Login(c *gin.Context) {

}

// @router /douyin/user/ [GET]
func GetUserInfo(c *gin.Context) {
	
}