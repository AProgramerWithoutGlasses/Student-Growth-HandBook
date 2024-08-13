package article

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	myErr "studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	utils "studentGrow/utils/readMessage"
	"studentGrow/utils/token"
)

// LikeController 点赞
func LikeController(c *gin.Context) {
	//解析数据
	json, err := utils.GetJsonvalue(c)
	if err != nil {
		zap.L().Error("Like() controller.article.GetJsonvalue err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	// 通过token获取username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("Like() controller.article.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	// 获取被点赞id
	id, err := json.GetInt("id")
	if err != nil {
		zap.L().Error("Like() controller.article.GetInt err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 获取被点赞类型
	likeType, err := json.GetInt("like_type")
	if err != nil {
		zap.L().Error("Like() controller.article.GetInt err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	//点赞
	err = article.LikeObjOrNot(strconv.Itoa(id), username, likeType)
	if err != nil {
		zap.L().Error("Like() controller.article.LikeObjOrNot err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	//返回数据
	res.ResponseSuccess(c, nil)
}

// CancelLikeController 取消点赞
//func CancelLikeController(c *gin.Context) {
//	//解析数据
//	j, e := utils.GetJsonvalue(c)
//	if e != nil {
//		fmt.Println("Like() controller.article.GetJsonvalue err=", e)
//		return
//	}
//
//	objId, userId, likeType, err := GetParam(j, "objId", "userId", "likeType")
//	if err != nil {
//		return
//	}
//
//	//取消点赞
//	err = article.CancelLike(objId, userId, likeType)
//	if err != nil {
//		fmt.Println("Like() controller.article.AnalyzeToMap err=", err)
//		return
//	}
//
//	//返回数据
//	res.ResponseSuccess(c, nil)
//}

// CheckLikeOrNotController 检查是否点赞
//func CheckLikeOrNotController(c *gin.Context) {
//	//解析数据
//	j, e := utils.GetJsonvalue(c)
//	if e != nil {
//		fmt.Println("Like() controller.article.GetJsonvalue err=", e)
//		return
//	}
//
//	objId, userId, likeType, err := GetParam(j, "objId", "userId", "likeType")
//	if err != nil {
//		fmt.Println("Like() controller.article.GetParam err=", err)
//		return
//	}
//
//	//检查用户是否已经点赞
//	result, err := redis.IsUserLiked(objId, userId, likeType)
//	if err != nil {
//		fmt.Println("Like() controller.article.IsUserLiked err=", err)
//		return
//	}
//
//	//返回数据
//	if result {
//		res.ResponseSuccess(c, "已点赞")
//	} else {
//		res.ResponseSuccess(c, "未点赞")
//	}
//}
//
//// GetObjLikeNumController 获取当前文章或评论的点赞数量
//func GetObjLikeNumController(c *gin.Context) {
//	//解析数据
//	j, e := utils.GetJsonvalue(c)
//	if e != nil {
//		fmt.Println("Like() controller.article.GetJsonvalue err=", e)
//		return
//	}
//	objId, _ := j.GetString("objId")
//	likeType, _ := j.GetInt("likeType")
//
//	//获取点赞数
//	likeSum, err := redis.GetObjLikes(objId, likeType)
//	if err != nil {
//		fmt.Println("Like() controller.article.GetObjLikes err=", err)
//		return
//	}
//	//返回数据
//	res.ResponseSuccess(c, likeSum)
//}

// 获取文章或评论的点赞集合
//func GetObjLikedUsersController() {
//
//}
