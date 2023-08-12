package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	model "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/user"
	service "github.com/czhi-bin/mini-tiktok-backend/biz/service/user"
)

// @router /douyin/user/ [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	
}

// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, model.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg: "Invalid parameters",
		})
		return
	}

	var user_id int64
	user_id, err = service.NewUserService(ctx, c).Register(&req)
	if err != nil {
		c.JSON(consts.StatusOK, model.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg: err.Error(),
		})
		return
	}

	// TODO: replace with proper response
	c.JSON(consts.StatusOK, model.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg: 	"Successfully registered",
		UserId: 	user_id, 		
		Token: 		"token",		// replace with proper token
	})
}

// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {

}

