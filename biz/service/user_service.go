package service

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/czhi-bin/mini-tiktok-backend/biz/dal/db"
	model "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/user"
	"github.com/czhi-bin/mini-tiktok-backend/pkg/utils"
)

type UserService struct {
	ctx context.Context
	c 	*app.RequestContext
}

// Creates a new user service
func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{
		ctx: ctx,
		c: c,
	}
}

func (s *UserService) Register(req *model.UserRegisterRequest) (user_id int64, err error) {
	user, err := db.QueryUser(req.Username)
	if err != nil {
		return -1, err
	}

	if *user != (db.User{}) {
		return -1, errors.New("User already exists")
	}

	hashedPassword, err := utils.Encrypt(req.Password)
	if err != nil {
		return -1, errors.New("Failed to encrypt password")
	}
	user_id, err = db.CreateUser(&db.User{
		UserName: 			req.Username,
		Password: 			hashedPassword,
		AvatarUrl: 			"default_avatar.jpg",
		BackgroundImageUrl: "default_background.jpg",
		Signature: 			"default_signature",
	})
	if err != nil {
		return -1, errors.New("Failed to create user")
	}

	return user_id, nil
}