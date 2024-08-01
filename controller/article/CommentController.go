package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	res "studentGrow/pkg/response"
	logic "studentGrow/service/comment"
	utils "studentGrow/utils/readMessage"
)

// PostCom 类型comment_type:‘article’or‘comment’;id;comment_content;comment_username
func PostCom(c *gin.Context) {
	//将数据解析到map中
	mp, err := utils.AnalyzeDataToMap(c)
	if err != nil {
		fmt.Println("PostCom() controller.article.AnalyzeToMap err=", err)
		return
	}
	//新增评论
	err = logic.PostComment(mp)

	if err != nil {
		fmt.Println("PostCom() controller.article.PostComment err=", err)
		return
	}
	res.ResponseSuccess(c, nil)
}
