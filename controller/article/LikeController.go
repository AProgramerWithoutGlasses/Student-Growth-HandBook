package article

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/gin-gonic/gin"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	utils "studentGrow/utils/readMessage"
)

func GetParm(j *jsonvalue.V, a, b, c string) (objId, userId string, likeType int, err error) {
	objId, err = j.GetString(a)
	if err != nil {
		fmt.Println("GetParm() controller.article.GetString err=", err)
		return objId, userId, likeType, err
	}
	userId, err = j.GetString(b)
	if err != nil {
		fmt.Println("GetParm() controller.article.GetString err=", err)
		return objId, userId, likeType, err
	}
	likeType, err = j.GetInt(c)
	if err != nil {
		fmt.Println("GetParm() controller.article.GetInt err=", err)
		return objId, userId, likeType, err
	}
	return objId, userId, likeType, err
}

// Like 点赞
func Like(c *gin.Context) {
	//解析数据
	d, e := utils.AnalyzeDataToMyData(c)
	if e != nil {
		fmt.Println("Like() controller.article.AnalyzeToMap err=", e)
		return
	}
	//获取jsonvalue对象
	j, err := d.GetJsonValue()
	if err != nil {
		fmt.Println("Like() controller.article.GetJsonValue err=", e)
		return
	}

	objId, userId, likeType, err := GetParm(j, "objId", "userId", "likeType")
	if err != nil {
		fmt.Println("Like() controller.article.GetParm err=", err)
		return
	}
	//点赞
	article.Like(objId, userId, likeType)

	//返回数据
	res.ResponseSuccess(c, nil)
}

// CancelLike 取消点赞
func CancelLike(c *gin.Context) {
	//解析数据
	d, e := utils.AnalyzeDataToMyData(c)
	if e != nil {
		fmt.Println("Like() controller.article.AnalyzeToMap err=", e)
		return
	}
	//获取jsonvalue对象
	j, err := d.GetJsonValue()
	if err != nil {
		fmt.Println("Like() controller.article.GetJsonValue err=", e)
		return
	}

	objId, _ := j.GetString("objId")
	userId, _ := j.GetString("userId")
	likeType, _ := j.GetInt("likeType")

	//取消点赞
	article.CancelLike(objId, userId, likeType)

	//返回数据
	res.ResponseSuccess(c, nil)
}

// CheckLikeOrNot 检查是否点赞
func CheckLikeOrNot(c *gin.Context) {
	//解析数据
	d, e := utils.AnalyzeDataToMyData(c)
	if e != nil {
		fmt.Println("Like() controller.article.AnalyzeToMap err=", e)
		return
	}
	//获取jsonvalue对象
	j, err := d.GetJsonValue()
	if err != nil {
		fmt.Println("Like() controller.article.GetJsonValue err=", e)
		return
	}
	objId, _ := j.GetString("objId")
	userId, _ := j.GetString("userId")
	likeType, _ := j.GetInt("likeType")

	//检查用户是否已经点赞
	result := article.IsUserLiked(objId, userId, likeType)

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
	d, e := utils.AnalyzeDataToMyData(c)
	if e != nil {
		fmt.Println("Like() controller.article.AnalyzeToMap err=", e)
		return
	}
	//获取jsonvalue对象
	j, err := d.GetJsonValue()
	if err != nil {
		fmt.Println("Like() controller.article.GetJsonValue err=", e)
		return
	}
	objId, _ := j.GetString("objId")
	likeType, _ := j.GetInt("likeType")

	//获取点赞数
	likeSum := article.GetObjLikes(objId, likeType)

	//返回数据
	res.ResponseSuccess(c, likeSum)
}
