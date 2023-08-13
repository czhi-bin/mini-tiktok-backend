package jwt

import (
	// "net/http"
	"strconv"
	"time"

	// "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	// "github.com/czhi-bin/mini-tiktok-backend/biz/model/common"
	// model "github.com/czhi-bin/mini-tiktok-backend/biz/model/basic/user"
)

func GenerateToken(UserId int64) (string, error) {
	claims := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iss": "mini-tiktok-backend",
		"aud": strconv.FormatInt(UserId, 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("dnekcba-kotkit-inim"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("dnekcba-kotkit-inim"), nil
	},	jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithIssuer("mini-tiktok-backend"))
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenUnverifiable
}

// func JWTMiddleWare() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var req model.UserLoginRequest
// 		err := c.BindQuery(&req)
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.JSON(http.StatusOK, common.CommonResponse{
// 				StatusCode: -1,
// 				StatusMsg: "Unauthorized (No token))",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		claims, err := ParseToken(tokenString)
// 		if err != nil {
// 			c.JSON(http.StatusOK, common.CommonResponse{
// 				StatusCode: -1,
// 				StatusMsg: "Invalid token (Invalid signature))",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		_, err = strconv.ParseInt(claims["aud"].(string), 10, 64)
// 		if err != nil {
// 			c.JSON(http.StatusOK, common.CommonResponse{
// 				StatusCode: -1,
// 				StatusMsg: "Invalid token (Invalid user id)",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		// if UserId != c. {
// 		// 	c.JSON(http.StatusOK, common.CommonResponse{
// 		// 		StatusCode: -1,
// 		// 		StatusMsg: "Invalid token (Invalid user id)",
// 		// 	})
// 		// 	c.Abort()
// 		// }
// 		c.Next()
// 	}
// }