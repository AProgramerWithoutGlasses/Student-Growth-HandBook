package stuManage

//	{ 响应
//		code:
//		msg:
//		data: {
//				year: [2021, 2022, 2023, 2024]
//				class: [计科221, 计科222, ]
//				stuInfo: [{name: "张四", userName: "20221544", password: "22mm", gender: true, }, {}, {}]
//				allStudentCount: 230
//			   }
//	}
//
//
//	{
//		year:
//		class:
//		gender:
//		isDisable:
//
//		searchSelect:
//		searchMessage:
//
//	}

type returnDataStruct struct {
	year    []string
	class   []string
	stuInfo []map[string]string
}

type resopseStruct struct {
	PlusTime string
	class    string
	gender   string
	Ban      string

	name        string
	username    string
	PhoneNumber string
}

//func QueryStu(c *gin.Context) {
//	queryStuReq, err := readMessage.GetJsonvalue(c)
//	if err != nil {
//		fmt.Println("queryStuController readMessage.GetJsonvalue() err :", err)
//	}
//
//	//notEmptyPullMap, notEmptyQueryMap := service.GetNotEmptyFieldMap(queryStuReq)
//
//}
