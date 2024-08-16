package models

import "github.com/dgrijalva/jwt-go"

// Login 后台登录的结构体
type Login struct {
	Username string `json:"username" binding:"len=11"`
	Password string `json:"password" binding:"max=17"`
	Code     string `json:"verify" binging:""`
	Id       string `json:"verifyId"`
}

// Claims Token结构体
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// TagAmount 文章tag和数量结构体
type TagAmount struct {
	Tag   int `json:"tag" gorm:"column:tag_id"`
	Count int
}

// StarBack 成长之星返回前端的结构体
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

// StarClass 成长之星按班级分类的结构体
type StarClass struct {
	ClassName string   `json:"className"`
	ClassStar []string `json:"classStar"`
}

// StarGrade 年级之星数据的结构体
type StarGrade struct {
	GradeName  string `json:"gradeName"`
	GradeClass string `json:"gradeClass"`
}

// 前台成长之星数据的结构体
type StarStu struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	UserHeadshot string `json:"user_headshot"`
}
