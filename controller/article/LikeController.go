package article

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/gin-gonic/gin"
	"studentGrow/dao/redis"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	utils "studentGrow/utils/readMessage"
)

func GetParam(j *jsonvalue.V, a, b, c string) (objId, userId string, likeType int, err error) {
	objId, err = j.GetString(a)
	if err != nil {
		fmt.Println("GetParam() controller.article.GetString err=", err)
		return objId, userId, likeType, err
	}
	userId, err = j.GetString(b)
	if err != nil {
		fmt.Println("GetParam() controller.article.GetString err=", err)
		return objId, userId, likeType, err
	}
	likeType, err = j.GetInt(c)
	if err != nil {
		fmt.Println("GetParam() controller.article.GetInt err=", err)
		return objId, userId, likeType, err
	}
	return objId, userId, likeType, err
}

// Like 点赞
func Like(c *gin.Context) {
	//解析数据
	j, e := utils.GetJsonvalue(c)
	if e != nil {
		fmt.Println("Like() controller.article.GetJsonvalue err=", e)
		return
	}

	objId, userId, likeType, err := GetParam(j, "objId", "userId", "likeType")
	if err != nil {
		fmt.Println("Like() controller.article.GetParam err=", err)
		return
	}
	//点赞
	err = article.Like(objId, userId, likeType)
	if err != nil {
		fmt.Println("Like() controller.article.Like err=", err)
		return
	}

	//返回数据
	res.ResponseSuccess(c, nil)
}

// CancelLike 取消点赞
func CancelLike(c *gin.Context) {
	//解析数据
	j, e := utils.GetJsonvalue(c)
	if e != nil {
		fmt.Println("Like() controller.article.GetJsonvalue err=", e)
		return
	}

	objId, userId, likeType, err := GetParam(j, "objId", "userId", "likeType")
	if err != nil {
		return
	}

	//取消点赞
	err = article.CancelLike(objId, userId, likeType)
	if err != nil {
		fmt.Println("Like() controller.article.AnalyzeToMap err=", err)
		return
	}

	//返回数据
	res.ResponseSuccess(c, nil)
}

// CheckLikeOrNot 检查是否点赞
func CheckLikeOrNot(c *gin.Context) {
	//解析数据
	j, e := utils.GetJsonvalue(c)
	if e != nil {
		fmt.Println("Like() controller.article.GetJsonvalue err=", e)
		return
	}

	objId, userId, likeType, err := GetParam(j, "objId", "userId", "likeType")
	if err != nil {
		fmt.Println("Like() controller.article.GetParam err=", err)
		return
	}

	//检查用户是否已经点赞
	result, err := redis.IsUserLiked(objId, userId, likeType)
	if err != nil {
		fmt.Println("Like() controller.article.IsUserLiked err=", err)
		return
	}

	//返回数据
	if result {
		res.ResponseSuccess(c, "已点赞")
	} else {
		res.ResponseSuccess(c, "未点赞")
	}
}

// GetObjLikeNum 获取当前文章或评论的点赞数量
func GetObjLikeNum(c *gin.Context) {
	//解析数据
	j, e := utils.GetJsonvalue(c)
	if e != nil {
		fmt.Println("Like() controller.article.GetJsonvalue err=", e)
		return
	}
	objId, _ := j.GetString("objId")
	likeType, _ := j.GetInt("likeType")

	//获取点赞数
	likeSum, err := redis.GetObjLikes(objId, likeType)
	if err != nil {
		fmt.Println("Like() controller.article.GetObjLikes err=", err)
		return
	}
	//返回数据
	res.ResponseSuccess(c, likeSum)
}
