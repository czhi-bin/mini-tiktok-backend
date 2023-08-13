package service

import (
	"github.com/gin-gonic/gin"

	model "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/feed"
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

func (s *FeedService) Feed(req *model.FeedRequest) (*model.FeedResponse, error) {
	// TODO: implement
	return &model.FeedResponse{
		VideoList: nil,
		NextTime:  -1,
	}, nil
}
