package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var SignSecret = 100

type UserToken struct {
	Name     string `gorm:"column:name" json:"name"`
	PassWord string `gorm:"column:password" json:"password"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(name, password string) (string, error) {
	// 创建一个我们自己的声明
	userToken := UserToken{
		Name:     name,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "YyNnn",                                    // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userToken)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(SignSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*UserToken, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserToken{}, func(token *jwt.Token) (i interface{}, err error) {
		return SignSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*UserToken); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		if len(authHeader) != 32 {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "报文头鉴权失败",
			})
			c.Abort()
			return
		}
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
