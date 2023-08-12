package db

import (
	"github.com/czhi-bin/mini-tiktok-backend/pkg/constants"
)

type User struct {
	ID	   				int64  `json:"id"`
	UserName  			string `json:"user_name"`
	Password  			string `json:"password"`
	AvatarUrl 			string `json:"avatar"`
	BackgroundImageUrl 	string `json:"background_image"`
	Signature 			string `json:"signature"`
}

func (User) TableName() string {
	return constants.UserTableName
}

func CreateUser(user *User) (int64, error) {
	err := DB.Create(user).Error
	if err != nil {
		return 0, err
	}
	
	return user.ID, nil
}

// QueryUser queries user by user_name
func QueryUser(userName string) (*User, error) {
	var user User
	err := DB.Where("user_name = ?", userName).Limit(1).Find(&user).Error
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}