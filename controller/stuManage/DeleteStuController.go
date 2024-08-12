package stuManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	"studentGrow/utils/readMessage"
)

func DeleteStuContro(c *gin.Context) {
	stuMessage, err := readMessage.GetJsonvalue(c)
	if err != nil {
		fmt.Println("readMessage.GetJsonvalue() err : ", err)
	}

	usernameValue, err := stuMessage.GetString("username")
	if err != nil {
		fmt.Println("username GetString() err : ", err)
	}

	// 根据学号获取id
	id, err := mysql.GetIdByUsername(usernameValue)
	if err != nil {
		fmt.Println("stuManage.DeleteStuContro() mysql.GetIdByUsername() err : ", err)
	}

	fmt.Println(id)
}
