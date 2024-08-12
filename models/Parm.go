package models

import "github.com/dgrijalva/jwt-go"

// 图形验证码
type Code struct {
	Id    string `json:"id"`
	Hcode string `json:"hcode"`
	B64   string `json:"b_64"`
}

// Login 后台登录的结构体
type Login struct {
	Username string `json:"username" binding:"len=11"`
	Password string `json:"password" binding:"max=17"`
	Code     string `json:"verify" binging:""`
	Id       string `json:"verifyId"`
	Level    string `json:"level"`
}

// Claims Token结构体
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// 文章tag和数量结构体
type TagAmount struct {
	Tag   string
	Count int
}

// 成长之星返回前端的结构体
type StarBack struct {
	Username           string `json:"username"`
	Frequency          int64  `json:"frequency"`
	Name               string `json:"name"`
	User_article_total int64  `json:"user_article_total"`
	Userfans           int64  `json:"userfans"`
	Score              int    `json:"score"`
	Hot                int    `json:"hot"`
	Status             bool   `json:"status"`
}
