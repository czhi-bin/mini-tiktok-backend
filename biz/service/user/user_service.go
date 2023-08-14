package service

import (
	"errors"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/czhi-bin/mini-tiktok-backend/biz/dal/db"
	"github.com/czhi-bin/mini-tiktok-backend/biz/middleware/jwt"
	userModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/user"
	commonModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/common"
	"github.com/czhi-bin/mini-tiktok-backend/pkg/utils"
)

type UserService struct {
	c *gin.Context
}

type UserAuth struct {
	UserId int64
	Token  string
}

// Creates a new user service
func NewService(c *gin.Context) *UserService {
	return &UserService{
		c: c,
	}
}

func (s *UserService) Register(req *userModel.UserRegisterRequest) (user_id int64, err error) {
	user, err := db.GetUserByName(req.Username)
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
		UserName:        req.Username,
		Password:        hashedPassword,
		Avatar:          "default_avatar.jpg",
		BackgroundImage: "default_background.jpg",
		Signature:       "default_signature",
	})
	if err != nil {
		return -1, errors.New("failed to create user")
	}

	return user_id, nil
}

func (s *UserService) Login(req *userModel.UserLoginRequest) (*UserAuth, error) {
	user, err := db.GetUserByName(req.Username)
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

	return &UserAuth{
		UserId: user.ID,
		Token:  token,
	}, nil
}

// GetUserInfo returns the user info of queryUserId according to the current user
func (s *UserService) GetUserInfo(req *userModel.UserRequest) (*commonModel.User, error) {
	queryUserId := req.UserId
	var currentUserId int64 = 0

	userInfo := &commonModel.User{Id: queryUserId}
	errChan := make(chan error, 7)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(7)

	go func() {
		// check if the user exists
		user, err := db.GetUserById(queryUserId)
		if err != nil {
			errChan <- err
		} else {
			userInfo.Name = user.UserName
			userInfo.Avatar = user.Avatar
			userInfo.BackgroundImage = user.BackgroundImage
			userInfo.Signature = user.Signature
		}
		wg.Done()
	}()

	go func() {
		// get the number of video published by the user
		workCount, err := db.GetWorkCount(queryUserId)
		if err != nil {
			errChan <- err
		} else {
			userInfo.WorkCount = workCount
		}
		wg.Done()
	}()

	go func() {
		// get the number of following of the user
		followingCount, err := db.GetFollowingCount(queryUserId)
		if err != nil {
			errChan <- err
		} else {
			userInfo.FollowCount = followingCount
		}
		wg.Done()
	}()

	go func() {
		// get the number of follower of the user
		followerCount, err := db.GetFollowerCount(queryUserId)
		if err != nil {
			errChan <- err
		} else {
			userInfo.FollowerCount = followerCount
		}
		wg.Done()
	}()

	go func() {
		// check if the current user follows the user
		// in this case, currentUser is the follower
		if currentUserId != 0 {
			// use for future use?...
			isFollowing, err := db.IsFollowing(queryUserId, currentUserId)
			if err != nil {
				errChan <- err
			} else {
				userInfo.IsFollow = isFollowing
			}
		} else {
			userInfo.IsFollow = false
		}
		wg.Done()
	}()

	go func() {
		// get the number of videos liked by the user
		favoriteCount, err := db.GetFavoriteCountByUserId(queryUserId)
		if err != nil {
			errChan <- err
		} else {
			userInfo.FavoriteCount = favoriteCount
		}
		wg.Done()
	}()

	go func() {
		// get the number of likes for video published by the user
		totalFavorited, err := db.GetTotalFavoritedByAuthorId(queryUserId)
		if err != nil {
			errChan <- err
		} else {
			userInfo.TotalFavorited = totalFavorited
		}
		wg.Done()
	}()

	wg.Wait()

	select {
	case err := <-errChan:
		return nil, err
	default:
	}

	return userInfo, nil
}
