package response

type Code int

// 常量初始化code值
const (
	SuccessCode     Code = 200
	ServerErrorCode Code = 501
)

// map用于存储每个code对应的提示信息
var codeMsgMap = map[Code]string{
	SuccessCode:     "操作成功",
	ServerErrorCode: "服务端错误",
}

// 用于获取code对应的提示信息
func (c Code) Msg() string {
	return codeMsgMap[c]
}
