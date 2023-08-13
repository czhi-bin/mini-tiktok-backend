package service

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/czhi-bin/mini-tiktok-backend/biz/dal/db"
	"github.com/czhi-bin/mini-tiktok-backend/biz/middleware/jwt"
	model "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/user"
	"github.com/czhi-bin/mini-tiktok-backend/biz/model/common"
	"github.com/czhi-bin/mini-tiktok-backend/pkg/utils"
)

type UserService struct {
	c 	*gin.Context
}

// Creates a new user service
func NewUserService(c *gin.Context) *UserService {
	return &UserService{
		c: c,
	}
}

func (s *UserService) Register(req *model.UserRegisterRequest) (user_id int64, err error) {
	user, err := db.QueryUser(req.Username)
	if err != nil {
		return -1, err
	}

	if *user != (db.User{}) {
		return -1, errors.New("user already exists")
	}

	hashedPassword, err := utils.Encrypt(req.Password)
	if err != nil {
		return -1, errors.New("failed to encrypt password")
	}

	user_id, err = db.CreateUser(&db.User{
		UserName: 			req.Username,
		Password: 			hashedPassword,
		Avatar: 			"default_avatar.jpg",
		BackgroundImage: 	"default_background.jpg",
		Signature: 			"default_signature",
	})
	if err != nil {
		return -1, errors.New("failed to create user")
	}

	return user_id, nil
}

func (s *UserService) Login(req *model.UserLoginRequest) (UserAuth *common.UserAuth, err error) {
	user, err := db.QueryUser(req.Username)
	if err != nil {
		return nil, err
	}

	if *user == (db.User{}) {
		return nil, errors.New("user does not exist")
	}

	if !utils.VerifyPassword(req.Password, user.Password) {
		return nil, errors.New("incorrect password")
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &common.UserAuth{
		UserId: 	user.ID,
		Token: 		token,
	}, nil
}