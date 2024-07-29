package models

import "github.com/dgrijalva/jwt-go"

// 图形验证码
type Code struct {
	Id  string
	B64 string
}

// Login 后台登录的结构体
type Login struct {
	Username string `json:"username" binding:"len=6"`
	Password string `json:"password" binding:"max=7"`
	Code     string `json:"code" binging:""`
	Id       string `json:"id"`
	Level    string `json:"level"`
}

// Token结构体
type Claims struct {
	Username string
	Password string
	jwt.StandardClaims
}
