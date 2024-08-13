package growth

// StarClass 班级管理员返回前端表格数据以选择
//func StarClass(c *gin.Context) {
//	var page struct {
//		page  int
//		limit int
//	}
//	err := c.ShouldBindQuery(&page)
//	token := c.GetHeader("token")
//	//获取username
//	username, err := token2.GetUsername(token)
//	if err != nil {
//		fmt.Println("starClass GetUsername err", err)
//		return
//	}
//	//查询班级
//	class, err := mysql.SelClass(username)
//	if err != nil {
//		fmt.Println("starClass SelClass err", err)
//		return
//	}
//	//查询班级成员的username
//	usernameslice, err := mysql.SelUsername(class)
//	if err != nil {
//		fmt.Println("StarClass SelUsername err", err)
//		return
//	}
//	starback, err := starService.StarGridClass(usernameslice)
//	if err != nil {
//		fmt.Println("StarClass starback err", err)
//		return
//	}
//
//	//换算页数
//
//}
