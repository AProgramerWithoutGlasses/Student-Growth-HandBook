package error

import (
	"fmt"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/errors"
	"github.com/gin-gonic/gin"
)
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
		Msg:  "record not found",
	}
}

// HasExistError 错误--数据内容重复冲突
func HasExistError() *Error {
	return &Error{
		Code: res.ServerErrorCode,
		Msg:  "data conflict",
	}
}

// DataFormatError 数据格式错误
func DataFormatError() *Error {
	return &Error{
		Code: res.ServerErrorCode,
		Msg:  "data format error",
	}
}

// CheckErrors 一键检查错误,并返回给客户端msg
func CheckErrors(err error, c *gin.Context) {
	if errors.Is(err, DataFormatError()) {
		// 前端发送数据格式错误
		res.ResponseErrorWithMsg(c, res.ServerErrorCode, DataFormatError().Msg)
		return
	}
	if errors.Is(err, HasExistError()) {
		// 数据已存在，发生冲突
		res.ResponseErrorWithMsg(c, res.ServerErrorCode, HasExistError().Msg)
		return
	}

	if errors.Is(err, NotFoundError()) {
		// 找不到对应数据
		res.ResponseErrorWithMsg(c, res.ServerErrorCode, NotFoundError().Msg)
		return
	}
	// 其他错误
	res.ResponseErrorWithMsg(c, res.ServerErrorCode, err.Error())
}
