package favorite

import (
	"net/http"

	"github.com/gin-gonic/gin"

	favoriteModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/interact/favorite"
	favoriteService "github.com/czhi-bin/mini-tiktok-backend/biz/service/favorite"
)

// @router /douyin/publish/action/ [POST]
func Action(c *gin.Context) {

}

// @router /douyin/publish/list/ [POST]
func List(c *gin.Context) {
	var err error
	var req favoriteModel.FavoriteListRequest
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, favoriteModel.FavoriteListResponse{
			StatusCode: -1,
			StatusMsg:  "Invalid parameters",
		})
		return
	}

	videos, err := favoriteService.NewService(c).List(&req)
	if err != nil {
		c.JSON(http.StatusOK, favoriteModel.FavoriteListResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, favoriteModel.FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "Favorite list retrieved successfully",
		VideoList: 	videos,
	})
}