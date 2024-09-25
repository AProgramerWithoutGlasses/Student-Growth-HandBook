package menuController

import (
	"github.com/gin-gonic/gin"
	"studentGrow/models"
	"studentGrow/pkg/response"
	service "studentGrow/service/permission"
	token2 "studentGrow/utils/token"
)

func MenuSide(c *gin.Context) {
	//定义返回前端的结构体
	var menusidar []models.Sidebar
	token := c.GetHeader("token")
	role, err := token2.GetRole(token)
	if err != nil {
		response.ResponseError(c, 400)
	}
	switch role {
	case "class":
		menusidar, err = service.MenuIdClass("class")
		if err != nil {
			response.ResponseError(c, 400)
			return
		}
	case "grade1", "grade2", "grade3", "grade4":
		menusidar, err = service.MenuIdClass("grade")
		if err != nil {
			response.ResponseError(c, 400)
			return
		}
	case "college":
		menusidar, err = service.MenuIdClass("college")
		if err != nil {
			response.ResponseError(c, 400)
			return
		}
	case "superman":
		menusidar, err = service.MenuIdClass("superman")
		if err != nil {
			response.ResponseError(c, 400)
			return
		}
	}
	response.ResponseSuccess(c, menusidar)
}
