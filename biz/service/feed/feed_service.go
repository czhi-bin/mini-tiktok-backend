package service

import (
	"github.com/gin-gonic/gin"

	feedModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/feed"
)

type FeedService struct {
	c *gin.Context
}

// Creates a new user service
func NewService(c *gin.Context) *FeedService {
	return &FeedService{
		c: c,
	}
}

func (s *FeedService) Feed(req *feedModel.FeedRequest) (*feedModel.FeedResponse, error) {
	// TODO: implement
	return &feedModel.FeedResponse{
		VideoList: nil,
		NextTime:  -1,
	}, nil
}
