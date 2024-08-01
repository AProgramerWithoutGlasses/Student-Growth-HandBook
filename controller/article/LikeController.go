package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	utils "studentGrow/utils/readMessage"
)

// Like 点赞
func Like(c *gin.Context) {
	//解析数据
	mp, e := utils.AnalyzeDataToMap(c)
	if e != nil {
		fmt.Println("Like() controller.article.AnalyzeToMap err=", e)
		return
	}
	objId := mp["objId"].(string)
	userId := mp["userId"].(string)
	likeType := mp["likeType"].(int)
	//点赞
	article.Like(objId, userId, likeType)

	//返回数据
	c.JSON(200, gin.H{"Msg": "OK"})
}

// CancelLike 取消点赞
func CancelLike(c *gin.Context) {
	//解析数据
	mp, e := utils.AnalyzeDataToMap(c)
	if e != nil {
		fmt.Println("CancelLike() controller.article.AnalyzeToMap err=", e)
		return
	}
	objId := mp["objId"].(string)
	userId := mp["userId"].(string)
	likeType := mp["likeType"].(int)

	//取消点赞
	article.CancelLike(objId, userId, likeType)

	//返回数据
	c.JSON(200, gin.H{"Msg": "OK"})
}

// CheckLikeOrNot 检查是否点赞
func CheckLikeOrNot(c *gin.Context) {
	//解析数据
	mp, e := utils.AnalyzeDataToMap(c)
	if e != nil {
		fmt.Println("CheckLikeOrNot() controller.article.AnalyzeToMap err=", e)
		return
	}
	objId := mp["objId"].(string)
	userId := mp["userId"].(string)
	likeType := mp["likeType"].(int)

	//检查用户是否已经点赞
	result := article.IsUserLiked(objId, userId, likeType)

	//返回数据
	if result {
		c.JSON(200, gin.H{"Msg": "OK"})
	} else {
		c.JSON(200, gin.H{"Msg": "NO"})
	}
}

// GetObjLikeNum 获取当前文章或评论的点赞数量
func GetObjLikeNum(c *gin.Context) {
	//解析数据
	mp, e := utils.AnalyzeDataToMap(c)
	if e != nil {
		fmt.Println("CheckLikeOrNot() controller.article.AnalyzeToMap err=", e)
		return
	}
	objId := mp["objId"].(string)
	likeType := mp["likeType"].(int)

	//获取点赞数
	likeSum := article.GetObjLikes(objId, likeType)

	//返回数据
	res.ResponseSuccess(c, likeSum)
}
