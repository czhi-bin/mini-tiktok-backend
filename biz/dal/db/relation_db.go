package db

import (
	"time"

	"gorm.io/gorm"

	"github.com/czhi-bin/mini-tiktok-backend/pkg/constants"
)

type Relation struct {
	ID        	int64     		`json:"id"`
	UserId    	int64     		`json:"user_id"`
	FollowerId  int64     		`json:"follow_id"`
	FollowedAt 	time.Time 		`json:"created_at"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleted_at"`
}

func (Relation) TableName() string {
	return constants.RelationTableName
}

// GetFollowingCount returns the number of users followed by the user
// In this case, the user is the follower
func GetFollowingCount(userId int64) (int64, error) {
	var count int64
	err := DB.Model(&Relation{}).Where("follower_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetFollowingByFollowerId returns the list of users (user ids) followed by the user
// In this case, the user is the follower
func GetFollowingsByFollowerId(followerId int64) ([]int64, error) {
	var relations []Relation
	err := DB.Where("follower_id = ?", followerId).Find(&relations).Error
	if err != nil {
		return nil, err
	}

	var users = make([]int64, 0, len(relations))
	for _, relation := range relations {
		users = append(users, relation.UserId)
	}

	return users, nil
}

// GetFollowerCount returns the number of users following the user
// In this case, the user is the followed
func GetFollowerCount(userId int64) (int64, error) {
	var count int64
	err := DB.Model(&Relation{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetFollowersByUserId returns the list of users (user ids) following the user
// In this case, the user is the followed
func GetFollowersByUserId(userId int64) ([]int64, error) {
	var relations []Relation
	err := DB.Where("user_id = ?", userId).Find(&relations).Error
	if err != nil {
		return nil, err
	}

	var users = make([]int64, 0, len(relations))
	for _, relation := range relations {
		users = append(users, relation.FollowerId)
	}

	return users, nil
}

// IsFollowing returns true if the follower is following user
func IsFollowing(userId, followerId int64) (bool, error) {
	var relation Relation
	err := DB.Where("user_id = ? AND follower_id = ?", userId, followerId).Limit(1).Find(&relation).Error
	if err != nil {
		return false, err
	}

	if relation == (Relation{}) {
		return false, nil
	}
	
	return true, nil
}
