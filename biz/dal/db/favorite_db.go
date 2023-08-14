package db

import (
	"time"

	"gorm.io/gorm"

	"github.com/czhi-bin/mini-tiktok-backend/pkg/constants"
)

type Favorite struct {
	ID	   		int64  			`json:"id"`
	UserId  	int64  			`json:"user_id"`
	VideoId	 	int64  			`json:"video_id"`
	LikedAt 	time.Time  		`json:"liked_at"`
	DeletedAt 	gorm.DeletedAt  `json:"unlike_at"`
}

func (Favorite) TableName() string {
	return constants.FavoriteTableName
}

// Returns the number of videos liked by a user
func GetFavoriteCountByUserId(userId int64) (int64, error) {
	var count int64
	err := DB.Model(&Favorite{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Returns the list of videos (video ids) liked by a user
func GetFavoriteIdList(userId int64) ([]int64, error) {
	var favoriteList []Favorite
	err := DB.Where("user_id = ?", userId).Find(&favoriteList).Error
	if err != nil {
		return nil, err
	}

	var videos = make([]int64, 0, len(favoriteList))
	for _, favorite := range favoriteList {
		videos = append(videos, favorite.VideoId)
	}

	return videos, nil
}

// Returns the total number of likes for videos publish by user
func GetTotalFavoritedByAuthorId(authorId int64) (int64, error) {
	var count int64
	err := DB.Table(constants.FavoriteTableName).Joins("JOIN videos ON videos.id = likes.video_id").
		Where("videos.author_id = ?", authorId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}