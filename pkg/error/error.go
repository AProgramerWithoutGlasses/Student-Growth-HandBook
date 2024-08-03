package error

import "fmt"
import res "studentGrow/pkg/response"

type Error struct {
	Code res.Code
	Msg  string
}

// 实现 Error 方法
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Msg)
}

// NotFoundError 错误--找不到相应数据
func NotFoundError() *Error {
	return &Error{
		Code: res.ServerErrorCode,
		Msg:  "not found",
	}
}

// HasExistError 错误--数据内容重复冲突
func HasExistError() *Error {
	return &Error{
		Code: res.ServerErrorCode,
		Msg:  "has existed",
	}
}
