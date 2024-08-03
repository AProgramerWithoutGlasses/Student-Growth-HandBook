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
	Code     string `json:"code" binging:""`
	Id       string `json:"id"`
	Level    string `json:"level"`
}

// Token结构体
type Claims struct {
	Username string
	Password string
	Role     string
	jwt.StandardClaims
}
