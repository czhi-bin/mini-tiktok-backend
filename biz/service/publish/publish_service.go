package service

import (
	"github.com/gin-gonic/gin"

	publishModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/publish"
	commonModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/common"
)

type FavoriteService struct {
	c *gin.Context
}

// Creates a new user service
func NewService(c *gin.Context) *FavoriteService {
	return &FavoriteService{
		c: c,
	}
}

// Publishes a new video
func (s *FavoriteService) Action(req *publishModel.PublishActionRequest) error {
	return nil
}

// Returns the list of videos published by the user
func (s *FavoriteService) List(req *publishModel.PublishListRequest) ([]*commonModel.Video, error) {
	// TODO: implement
	return nil, nil
}
