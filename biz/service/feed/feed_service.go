package service

import (
	"github.com/gin-gonic/gin"

	feedModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/feed"
	commonModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/common"
)

type FeedService struct {
	c *gin.Context
}

type Feed struct {
	VideoList []*commonModel.Video
	NextTime  int64
}

// Creates a new user service
func NewService(c *gin.Context) *FeedService {
	return &FeedService{
		c: c,
	}
}

func (s *FeedService) Feed(req *feedModel.FeedRequest) (*Feed, error) {
	// TODO: implement
	return &Feed{
		VideoList: nil,
		NextTime:  -1,
	}, nil
}
