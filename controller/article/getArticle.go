package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/logic"
	"studentGrow/pkg"
	utils "studentGrow/utils/readMessage"
)

func GetArticleId(c *gin.Context) {
	// 从c.Request.Body读取请求数据
	b, _ := c.GetRawData()

	//将数据解析到map中
	mp := utils.AnalyzeToMap(b)

	//查询结果
	err, user, article := logic.GetArticleLogic(mp)

	if err != nil {
		fmt.Println("GetArticleId() controller.article.getArticle.GetArticleLogic err=", err)
		c.JSON(501, pkg.ArticlePkg{
			Type: "getArticle",
			Data: nil,
			Msg:  "未找到该用户",
		})
		return
	}
	c.JSON(200, pkg.ArticlePkg{
		Type: "getArticle",
		Data: map[string]any{"user": user, "article": article},
		Msg:  "",
	})

}
