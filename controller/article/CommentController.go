package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	myErr "studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	logic "studentGrow/service/comment"
	utils "studentGrow/utils/readMessage"
	"studentGrow/utils/token"
)

// PostCom 发布评论
func PostCom(c *gin.Context) {
	// 读取前端数据
	json, err := utils.GetJsonvalue(c)
	if err != nil {
		fmt.Println("PostCom() controller.article.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	// 通过token获取username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		fmt.Println("PostCom() controller.article.GetUsername err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	//获取数据
	commentType, err := json.GetString("comment_type")
	if err != nil {
		fmt.Println("PostCom() controller.article.GetString err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	commentContent, err := json.GetString("comment_content")
	if err != nil {
		fmt.Println("PostCom() controller.article.GetString err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	id, err := json.GetInt("id")
	if err != nil {
		fmt.Println("PostCom() controller.article.GetString err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	//新增评论
	err = logic.PostComment(commentType, username, commentContent, id)

	if err != nil {
		fmt.Println("PostCom() controller.article.PostComment err=", err)
		return
	}
	res.ResponseSuccess(c, nil)
}

// GetLel1CommentsController 获取一级评论
func GetLel1CommentsController(c *gin.Context) {
	// 读取前端数据
	json, err := utils.GetJsonvalue(c)
	if err != nil {
		fmt.Println("GetLel1CommentsController() controller.article.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	// 获取文章id
	aid, err := json.GetInt("article_id")
	if err != nil {
		fmt.Println("GetLel1CommentsController() controller.article.GetInt err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	comments, err := mysql.QueryLevelOneComments(aid)
	if err != nil {
		fmt.Println("GetLel1CommentsController() controller.article.QueryLevelOneComments err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, comments)
}

// GetLe2CommentsController 获取二级评论
func GetLe2CommentsController(c *gin.Context) {
	// 读取前端数据
	json, err := utils.GetJsonvalue(c)
	if err != nil {
		fmt.Println("GetLe2CommentsController() controller.article.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	// 获取父级评论id
	pid, err := json.GetInt("comment_id")
	if err != nil {
		fmt.Println("GetLe2CommentsController() controller.article.GetInt err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	comments, err := mysql.QueryLevelTwoComments(pid)
	if err != nil {
		fmt.Println("GetLe2CommentsController() controller.article.QueryLevelTwoComments err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	res.ResponseSuccess(c, comments)
}

// DeleteCommentController 删除评论
func DeleteCommentController(c *gin.Context) {
	// 读取前端数据
	json, err := utils.GetJsonvalue(c)
	if err != nil {
		fmt.Println("DeleteCommentController() controller.article.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	// 获取评论id
	cid, err := json.GetInt("comment_id")
	if err != nil {
		fmt.Println("DeleteCommentController() controller.article.GetInt err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	// 删除评论
	err = mysql.DeleteComment(cid)
	if err != nil {
		fmt.Println("DeleteCommentController() controller.article.DeleteComment err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	res.ResponseSuccess(c, nil)
}
