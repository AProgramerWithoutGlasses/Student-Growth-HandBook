package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	"studentGrow/dao/redis"
	"studentGrow/models"
	pkg "studentGrow/pkg/response"
	"studentGrow/service/userService"
	"studentGrow/utils/token"
)

// RidCode 图形验证码返回前端
func RidCode(c *gin.Context) {
	//获取验证码
	id, b64, hcode, _ := userService.CaptchaGenerate()
	if id == "" || b64 == " " {
		pkg.ResponseError(c, 400)
		fmt.Println("RidCode  Login() err")
		return
	}
	code := map[string]any{
		"Id":    id,
		"Hcode": hcode,
		"B64":   b64,
	}
	pkg.ResponseSuccess(c, code)
}

// HLogin 登录验证，成功返回token
func HLogin(c *gin.Context) {
	// 定义用户实例
	var user = new(models.Login)
	//获取前端返回的数据
	if err := c.BindJSON(&user); err != nil {
		pkg.ResponseErrorWithMsg(c, 400, "Hlogin 获取数据失败")
		fmt.Println("Hlogin BindJSON(&userService) err")
		return
	}
	//验证密码
	if ok := userService.BVerify(user.Username, user.Password); !ok {
		pkg.ResponseErrorWithMsg(c, 400, "密码错误")
		return
	}
	//验证验证码
	if ok := userService.GetCodeAnswer(user.Id, user.Code); !ok {
		pkg.ResponseErrorWithMsg(c, 400, "验证码错误")
		return
	}
	//验证用户是否存在
	if ok := userService.BVerifyExit(user.Username); !ok {
		pkg.ResponseErrorWithMsg(c, 400, "用户不存在")
		return
	}
	//记录用户登录
	redis.UpdateVictor(user.Username)
	//查询用户角色
	casbinId, _ := mysql.SelCasId(user.Username)
	role, err := mysql.SelRole(casbinId)
	//获取生成的token
	tokenString, err := token.ReleaseToken(user.Username, user.Password, role)
	if err != nil {
		fmt.Println("Hlogin的login.ReleaseToken()")
		return
	}
	if err != nil {
		fmt.Println("HLogin SelId err")
		return
	}
	slice := map[string]any{
		"username": user.Username,
		"token":    tokenString,
		"role":     role,
	}
	//发送给前端
	pkg.ResponseSuccess(c, slice)
}

// QLogin 登录验证，成功返回token
func QLogin(c *gin.Context) {
	// 定义用户实例
	var user = new(models.Login)
	var role string
	//获取前端返回的数据
	if err := c.BindJSON(&user); err != nil {
		pkg.ResponseErrorWithMsg(c, 400, "Hlogin 获取数据失败")
		fmt.Println("Hlogin BindJSON(&userService) err")
		return
	}
	//验证密码
	if ok := userService.BVerify(user.Username, user.Password); !ok {
		pkg.ResponseErrorWithMsg(c, 400, "密码错误")
		return
	}
	//验证验证码
	if ok := userService.GetCodeAnswer(user.Id, user.Code); !ok {
		pkg.ResponseErrorWithMsg(c, 400, "验证码错误")
		return
	}
	//验证用户是否存在
	ok := userService.BVerifyExit(user.Username)
	if ok {
		role = "1"
	} else {
		role = "0"
	}
	//记录用户登录
	redis.UpdateVictor(user.Username)
	//获取生成的token
	tokenString, err := token.ReleaseToken(user.Username, user.Password, role)
	if err != nil {
		fmt.Println("Hlogin的login.ReleaseToken()")
		return
	}
	if err != nil {
		fmt.Println("HLogin SelId err")
		return
	}
	slice := map[string]any{
		"username": user.Username,
		"token":    tokenString,
		"role":     role,
	}
	//发送给前端
	pkg.ResponseSuccess(c, slice)
}
