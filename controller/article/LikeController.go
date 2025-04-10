package article

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	myErr "studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	"studentGrow/utils/token"
)

// LikeController 点赞
func LikeController(c *gin.Context) {

	in := struct {
		Id          int    `json:"id"`
		LikeType    int    `json:"like_type"`
		TarUsername string `json:"tar_username"`
	}{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		return
	}

	// 通过token获取username
	aToken := token.NewToken(c)
	user, exist := aToken.GetUser()
	if !exist {
		res.ResponseError(c, res.TokenError)
		zap.L().Error("token错误")
		return
	}

	username := user.Username

	//点赞
	err = article.LikeObjOrNot(strconv.Itoa(in.Id), username, in.TarUsername, in.LikeType)
	if err != nil {
		zap.L().Error("Like() controller.article.LikeObjOrNot err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	//返回数据
	res.ResponseSuccess(c, struct{}{})
}
