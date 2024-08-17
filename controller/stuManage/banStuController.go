package stuManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

func BanUserControl(c *gin.Context) {
	// 接收数据
	var user gorm_model.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("stuManage.BanStuControl() c.Bind() err : ", err)
		response.ResponseErrorWithMsg(c, 500, "stuManage.BanStuControl() c.Bind() failed : "+err.Error())
		return
	}

	name, temp, err := service.BanUserService(user.Username)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "stuManage.BanStuControl() service.BanUserService() failed : "+err.Error())
		return
	}

	// 响应
	if temp == 0 {
		response.ResponseSuccess(c, "已将用户"+name+"封禁")
	} else if temp == 1 {
		response.ResponseSuccess(c, "已将用户"+name+"取消封禁")
	}

}
