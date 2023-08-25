package publish

import (
	"net/http"

	"github.com/gin-gonic/gin"

	publishModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/publish"
	publishService "github.com/czhi-bin/mini-tiktok-backend/biz/service/publish"
)

// @router /douyin/publish/action/ [POST]
func Action(c *gin.Context) {
	var err error
	var req publishModel.PublishActionRequest
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, publishModel.PublishActionResponse{
			StatusCode: -1,
			StatusMsg:  "Invalid parameters",
		})
		return
	}

	err = publishService.NewService(c).Action(&req)
}

// @router /douyin/publish/list/ [POST]
func List(c *gin.Context) {
	var err error
	var req publishModel.PublishListRequest
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, publishModel.PublishListResponse{
			StatusCode: -1,
			StatusMsg:  "Invalid parameters",
		})
		return
	}

	videos, err := publishService.NewService(c).List(&req)
	if err != nil {
		c.JSON(http.StatusOK, publishModel.PublishListResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, publishModel.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "List of published work retrieved successfully",
		VideoList: 	videos,
	})
}