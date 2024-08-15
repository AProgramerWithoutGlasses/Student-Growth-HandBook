package growth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	"studentGrow/service/starService"
	token2 "studentGrow/utils/token"
)

// Search 搜索表格数据
func Search(c *gin.Context) {
	var usernamesli []string
	//获取前端传来的数据
	var datas struct {
		Name  string `form:"search"`
		Page  int    `form:"page"`
		Limit int    `form:"pageCapacity"`
	}
	err := c.Bind(&datas)
	if err != nil {
		zap.L().Error("Search Bind err", zap.Error(err))
		response.ResponseError(c, response.ServerErrorCode)
		return
	}

	//得到登录者的角色和账号
	token := c.GetHeader("token")
	role, err := token2.GetRole(token)
	username, err := token2.GetUsername(token)
	if err != nil {
		zap.L().Error("Search err", zap.Error(err))
		response.ResponseError(c, response.ServerErrorCode)
		return
	}

	//查找stars库中所有的数据
	//查找权限下的数据
	alluser, err := mysql.SelStarUser()
	if err != nil {
		fmt.Println("StarGrade err", err)
		response.ResponseErrorWithMsg(c, 400, "获取表格数据失败")
	}

	//根据角色分类
	switch role {
	case "class":
		//查找班级
		class, err := mysql.SelClass(username)
		if err != nil {
			zap.L().Error("Search SelClass err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
		if datas.Name == "" {
			usernamesli, err = mysql.SelUsername(class)

		} else {
			usernamesli, err = mysql.SelSearchUser(datas.Name, class)
		}

		if err != nil {
			zap.L().Error("Search SelSearchUser err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
	case "grade1":
		if datas.Name == "" {
			usernamesli, err = starService.StarGuidGrade(alluser, 1)
		} else {
			usernamesli, err = starService.SearchGrade(datas.Name, 1)
			if err != nil {
				zap.L().Error("Search GetEnrollmentYear err", zap.Error(err))
				response.ResponseError(c, response.ServerErrorCode)
				return
			}
		}
	case "grade2":
		if datas.Name == "" {
			usernamesli, err = starService.StarGuidGrade(alluser, 2)
		} else {
			usernamesli, err = starService.SearchGrade(datas.Name, 2)
			if err != nil {
				zap.L().Error("Search GetEnrollmentYear err", zap.Error(err))
				response.ResponseError(c, response.ServerErrorCode)
				return
			}
		}
	case "grade3":
		if datas.Name == "" {
			usernamesli, err = starService.StarGuidGrade(alluser, 3)
		} else {
			usernamesli, err = starService.SearchGrade(datas.Name, 3)
			if err != nil {
				zap.L().Error("Search GetEnrollmentYear err", zap.Error(err))
				response.ResponseError(c, response.ServerErrorCode)
				return
			}
		}
	case "grade4":
		if datas.Name == "" {
			usernamesli, err = starService.StarGuidGrade(alluser, 4)
		} else {
			usernamesli, err = starService.SearchGrade(datas.Name, 4)
			if err != nil {
				zap.L().Error("Search GetEnrollmentYear err", zap.Error(err))
				response.ResponseError(c, response.ServerErrorCode)
				return
			}
		}
	case "college":
		if datas.Name == "" {
			usernamesli, err = mysql.SelStarColl()
		} else {
			usernamesli, err = mysql.SelSearchColl(datas.Name)
		}

		if err != nil {
			zap.L().Error("Search SelSearchColl err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
	case "superman":
		if datas.Name == "" {
			usernamesli, err = mysql.SelStarColl()
		} else {
			usernamesli, err = mysql.SelSearchColl(datas.Name)
		}

		if err != nil {
			zap.L().Error("Search SelSearchColl err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
	}

	//判断是否找到数据
	if usernamesli == nil {
		response.ResponseErrorWithMsg(c, 400, "数据未找到")
		return
	}

	//表格所需所有数据
	starback, err := starService.StarGrid(usernamesli)
	if err != nil {
		fmt.Println("StarGrade starback err", err)
		return
	}
	total := len(starback)
	//实现分页
	tableData := starService.PageQuery(starback, datas.Page, datas.Limit)
	data := map[string]any{
		"tableData": tableData,
		"total":     total,
	}
	response.ResponseSuccess(c, data)
}
