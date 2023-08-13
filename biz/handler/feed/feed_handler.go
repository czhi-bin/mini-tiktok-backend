package feed

import (
	"net/http"

	"github.com/gin-gonic/gin"

	feedModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/feed"
	"github.com/czhi-bin/mini-tiktok-backend/biz/model/common"
	feedService "github.com/czhi-bin/mini-tiktok-backend/biz/service/feed"
)

// @router /douyin/feed/ [GET]
func Feed(c *gin.Context) {
	var err error
	var req feedModel.FeedRequest
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, feedModel.FeedResponse{
			CommonResponse: &common.CommonResponse{
				StatusCode: -1,
				StatusMsg:  "Invalid parameters",
			},
		})
		return
	}

	feed, err := feedService.NewService(c).Feed(&req)
	if err != nil {
		c.JSON(http.StatusOK, feedModel.FeedResponse{
			CommonResponse: &common.CommonResponse{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, feedModel.FeedResponse{
		CommonResponse: &common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "Feed retrieved successfully",
		},
		VideoList: feed.VideoList,
		NextTime:  feed.NextTime,
	})
}
