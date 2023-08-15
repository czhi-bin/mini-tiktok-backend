package service

import (
	"github.com/gin-gonic/gin"

	favoriteModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/interact/favorite"
	commonModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/common"
)

type FavoriteService struct {
	c *gin.Context
}

type Favorite struct {

}

// Creates a new user service
func NewService(c *gin.Context) *FavoriteService {
	return &FavoriteService{
		c: c,
	}
}

// Returns the list of videos that the user has liked
func (s *FavoriteService) List(req *favoriteModel.FavoriteListRequest) ([]*commonModel.Video, error) {
	// TODO: implement
	return nil, nil
}
