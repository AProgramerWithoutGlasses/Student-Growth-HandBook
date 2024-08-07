package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"studentGrow/models"
)

// 中间件检验token是否合法
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取前端传过来的信息
		tokenString := c.GetHeader("token")
		fmt.Print("请求token", tokenString)
		//验证前端传过来的token格式，不为空，开头为Bearer
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(400, gin.H{"code": "400", "msg": "验证失败"})
			c.Abort()
			return
		}
		//验证通过，提取有效部分（除去Bearer)
		tokenString = tokenString[7:] //截取字符
		//解析token
		token, _, err := ParseToken(tokenString)
		//解析失败||解析后的token无效
		if err != nil || !token.Valid {
			c.JSON(400, gin.H{"code": "400", "msg": "验证失败"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// ParseToken 解析从前端获取到的token值
func ParseToken(tokenString string) (*jwt.Token, *models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	return token, claims, err
}

//通过token获取username

func GetUsername(tokenString string) (string, error) {
	tokenString = tokenString[7:]
	_, claims, err := ParseToken(tokenString)
	if err != nil {
		fmt.Println("GetUsername  ParseToken() err:", err.Error())
		return "", err
	}
	return claims.Username, nil
}
