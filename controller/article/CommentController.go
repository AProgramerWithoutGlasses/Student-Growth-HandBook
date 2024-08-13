package article

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	myErr "studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	"studentGrow/service/comment"
	utils "studentGrow/utils/readMessage"
	"studentGrow/utils/token"
)

// PostCom 发布评论
func PostCom(c *gin.Context) {
	// 读取前端数据
	json, err := utils.GetJsonvalue(c)
	if err != nil {
		zap.L().Error("PostCom() controller.article.GetJsonvalue err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	// 通过token获取username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("PostCom() controller.article.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	//获取数据
	commentType, err := json.GetString("comment_type")
	if err != nil {
		zap.L().Error("PostCom() controller.article.GetString err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	commentContent, err := json.GetString("comment_content")
	if err != nil {
		zap.L().Error("PostCom() controller.article.GetString err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	id, err := json.GetInt("id")
	if err != nil {
		zap.L().Error("PostCom() controller.article.GetInt err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	//新增评论
	err = comment.PostComment(commentType, username, commentContent, id)

	if err != nil {
		zap.L().Error("PostCom() controller.article.PostComment err=", zap.Error(err))
		return
	}
	res.ResponseSuccess(c, nil)
}

// GetLel1CommentsController 获取一级评论
func GetLel1CommentsController(c *gin.Context) {
	var input struct {
		Aid      int    `json:"article_id"`
		SortWay  string `json:"comment_way"`
		Limit    int    `json:"comment_count"`
		Page     int    `json:"comment_page"`
		Username string `json:"username"`
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		zap.L().Error("GetLel1CommentsController() controller.article.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	comments, err := comment.GetLel1CommentsService(input.Aid, input.Limit, input.Page, input.Username, input.SortWay)
	if err != nil {
		zap.L().Error("GetLel1CommentsController() controller.article.GetLel1CommentsService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 获取文章评论数
	commentNum, err := mysql.QueryArticleCommentNum(input.Aid)
	if err != nil {
		zap.L().Error("GetLel1CommentsController() controller.article.QueryArticleCommentNum err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, map[string]any{
		"comment_list": comments,
		"comment_num":  commentNum,
	})
}

// GetSonCommentsController 获取子评论列表
func GetSonCommentsController(c *gin.Context) {
	var input struct {
		Cid      int    `json:"comment_id"`
		Username string `json:"username"`
		Limit    int    `json:"limit"`
		Page     int    `json:"page"`
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		zap.L().Error("GetSonCommentsController() controller.article.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	comments, err := comment.GetLelSonCommentListService(input.Cid, input.Limit, input.Page, input.Username)

	if err != nil {
		zap.L().Error("GetSonCommentsController() controller.article.GetLelSonCommentListService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	res.ResponseSuccess(c, map[string]any{
		"comment_se_list": comments,
	})
}

// DeleteCommentController 删除评论
func DeleteCommentController(c *gin.Context) {
	// 读取前端数据
	json, err := utils.GetJsonvalue(c)
	if err != nil {
		zap.L().Error("DeleteCommentController() controller.article.GetJsonvalue err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 获取评论id
	cid, err := json.GetInt("comment_id")
	if err != nil {
		zap.L().Error("DeleteCommentController() controller.article.GetInt err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	// 删除评论
	err = mysql.DeleteComment(cid)
	if err != nil {
		zap.L().Error("DeleteCommentController() controller.article.DeleteComment err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	res.ResponseSuccess(c, nil)
}
