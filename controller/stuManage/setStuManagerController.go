package stuManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/pkg/response"
	"studentGrow/utils/readMessage"
)

func SetStuManagerControl(c *gin.Context) {
	reqMessage, err := readMessage.GetJsonvalue(c)
	if err != nil {
		fmt.Println("get json value error:", err)
		response.ResponseErrorWithMsg(c, 400, err.Error())
	}

	selectedStuArr, err := reqMessage.GetArray("selected_students")
	if err != nil {
		fmt.Println("stuManage.stuManagerControl() reqMessage.GetArray():", err)
		response.ResponseErrorWithMsg(c, 400, err.Error())
	}
	fmt.Println(selectedStuArr)
}
