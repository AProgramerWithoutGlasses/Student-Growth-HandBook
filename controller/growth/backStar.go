package growth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	"studentGrow/service/starService"
	token2 "studentGrow/utils/token"
)

// StarClass 班级管理员返回前端表格数据以选择
func StarClass(c *gin.Context) {
	var page struct {
		Page  int `form:"page"`
		Limit int `form:"limit"`
	}
	err := c.Bind(&page)
	token := c.GetHeader("token")
	//获取username
	username, err := token2.GetUsername(token)
	if err != nil {
		fmt.Println("starClass GetUsername err", err)
		return
	}
	//查询班级
	class, err := mysql.SelClass(username)
	if err != nil {
		fmt.Println("starClass SelClass err", err)
		return
	}
	//查询班级成员的username
	usernameslice, err := mysql.SelUsername(class)
	if err != nil {
		fmt.Println("StarClass SelUsername err", err)
		return
	}
	starback, err := starService.StarGrid(usernameslice)
	if err != nil {
		fmt.Println("StarClass starback err", err)
		return
	}

	//换算页数
	data := starService.PageQuery(starback, page.Page, page.Limit)
	response.ResponseSuccess(c, data)
}

// StarGrade 年级管理员返回前端表格以选择
func StarGrade(c *gin.Context) {
	//从前端获取数据
	var usernameslice []string
	var page struct {
		Page  int `form:"page"`
		Limit int `form:"limit"`
	}
	err := c.Bind(&page)
	if err != nil {
		fmt.Println("StarGrade err", err)
		response.ResponseErrorWithMsg(c, 400, "Bind接收数据失败")
	}
	//获取角色
	token := c.GetHeader("token")
	role, err := token2.GetRole(token)
	if err != nil {
		fmt.Println("StarGrade GetRole err", err)
		response.ResponseError(c, 400)
		return
	}
	//查找权限下的数据
	alluser, err := mysql.SelStarUser()
	if err != nil {
		fmt.Println("StarGrade err", err)
		response.ResponseErrorWithMsg(c, 400, "获取表格数据失败")
	}

	//查找对应角色对应的数据
	switch role {
	case "grade1":
		usernameslice, err = starService.StarGuidGrade(alluser, 1)
		if err != nil {
			fmt.Println("StarGrade grade1 StarGuidGrade err", err)
			return
		}
	case "grade2":
		usernameslice, err = starService.StarGuidGrade(alluser, 2)
		if err != nil {
			fmt.Println("StarGrade grade2 StarGuidGrade err", err)
			return
		}
	case "grade3":
		usernameslice, err = starService.StarGuidGrade(alluser, 3)
		if err != nil {
			fmt.Println("StarGrade grade3 StarGuidGrade err", err)
			return
		}
	case "grade4":
		usernameslice, err = starService.StarGuidGrade(alluser, 4)
		if err != nil {
			fmt.Println("StarGrade grade4 StarGuidGrade err", err)
			return
		}
	}

	//表格所需所有数据
	starback, err := starService.StarGrid(usernameslice)
	if err != nil {
		fmt.Println("StarGrade starback err", err)
		return
	}

	//实现分页
	data := starService.PageQuery(starback, page.Page, page.Limit)
	response.ResponseSuccess(c, data)
}

// StarCollege 院级管理员及超级管理员返回前端表格
func StarCollege(c *gin.Context) {
	//从前端获取页数及容量
	var page struct {
		Page  int `form:"page"`
		Limit int `form:"limit"`
	}
	err := c.Bind(&page)
	if err != nil {
		fmt.Println("StarCollege err", err)
		response.ResponseErrorWithMsg(c, 400, "Bind接收数据失败")
	}

	//获取学号合集
	userslice, err := mysql.SelStarColl()
	if err != nil {
		fmt.Println("StarCollege SelStarColl err", err)
		response.ResponseErrorWithMsg(c, 400, "院级管理员获取数据失败")
	}

	//表格所需所有数据
	starback, err := starService.StarGrid(userslice)
	if err != nil {
		fmt.Println("StarCollege starback err", err)
		return
	}

	//实现分页
	data := starService.PageQuery(starback, page.Page, page.Limit)
	response.ResponseSuccess(c, data)
}

// 搜索表格数据
func Search() {

}
