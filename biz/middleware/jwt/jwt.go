package jwt

import (
	"net/http"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/czhi-bin/mini-tiktok-backend/biz/dal/db"
	userModel "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/user"
	"github.com/czhi-bin/mini-tiktok-backend/pkg/utils"
)

var (
	JWTMiddleWare 	*jwt.GinJWTMiddleware
	identityKey 	= "user_id"
)

func Init() {
	var err error
	JWTMiddleWare, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:	   			"mini-tiktok-backend",
		SigningAlgorithm: 	"HS256",
		Key: 				[]byte("dnekcba-kotkit-inim"),
		Timeout: 			time.Hour,
		MaxRefresh: 		time.Hour * 24,
		IdentityKey: 		identityKey,
		TokenLookup: 		"query:token, form:token",

		// To verify password with hashedPassword in DB
		// Returns UserId as the verifier
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var req userModel.UserLoginRequest
			if err := c.Bind(&req); err != nil {
				return nil, err
			}
			
			user, err := db.GetUserByName(req.Username)
			if err != nil {
				return nil, err
			}
			if ok := utils.VerifyPassword(req.Password, user.Password); !ok {
				return nil, jwt.ErrFailedAuthentication
			}

			c.Set("user_id", user.ID)
			return user.ID, nil
		},

		// Sets the payload in the token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					"aud": v,
				}
			}
			return jwt.MapClaims{}
		},

		// Build login response if password is verified correctly
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.Set("token", token)
			c.Next()
		},

		// Verify token and get the id of the logged-in user
		Authorizator: func(data interface{}, c *gin.Context) bool {
			fmt.Println("AUTHORIZATOR CALLED")
			fmt.Println("data: ", data)
			if v, ok := data.(int64); ok {
				currentUserId := int64(v)
				c.Set("current_user_id", currentUserId)
				return true
			}
			return false
		},

		// Validation failed
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusOK, userModel.UserLoginResponse{
				StatusCode: -1,
				StatusMsg: message,
			})
		},
	})
	
	if err != nil {
		panic("JWT init Error: " + err.Error())
	}
}