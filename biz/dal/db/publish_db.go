package db

import (
	"time"

	"github.com/czhi-bin/mini-tiktok-backend/pkg/constants"
)

type Video struct {
	ID				int64  		`json:"id"`
	AuthorId		int64  		`json:"author_id"`
	VideoUrl		string  	`json:"video_url"`
	CoverUrl		string  	`json:"video_cover_url"`
	Title			string  	`json:"title"`
	PublishTime 	time.Time 	`json:"publish_time"`
}

func (Video) TableName() string {
	return constants.VideoTableName
}

// Returns the number of video published by the user
func GetWorkCount(userId int64) (int64, error) {
	var count int64
	err := DB.Model(&Video{}).Where("author_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Returns the list of video published by the user
func GetVideoByAuthorId(authorId int64) ([]Video, error) {
	var videos []Video
	err := DB.Where("author_id = ?", authorId).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	return videos, nil
}