package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 服务端密钥
var jwtSecret = []byte("setting.JwtSecret")

type Claims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(username, user_id string) (string, error) {
	nowTime := time.Now()
	// 设置过期时间为 3 小时
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		user_id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
