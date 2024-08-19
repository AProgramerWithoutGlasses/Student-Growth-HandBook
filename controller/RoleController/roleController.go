package RoleController

import (
	"github.com/gin-gonic/gin"
	"studentGrow/pkg/response"
	service "studentGrow/service/permission"
)

func RoleList(c *gin.Context) {
	rolelist, err := service.RoleData()
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	response.ResponseSuccess(c, rolelist)
}
