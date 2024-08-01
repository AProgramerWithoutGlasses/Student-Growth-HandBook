package comment

import (
	"fmt"
	"studentGrow/dao/mysql"
	"studentGrow/utils/readMessage"
)

// PostComment 发布评论
func PostComment(m map[string]any) (err error) {
	//类型comment_type:‘article’or‘comment’;id;comment_content;comment_username

	//获取数据
	commentType := m["comment_type"].(string)
	commentContent := m["comment_content"].(string)
	commentUsername := m["comment_username"].(string)
	//获取用户id
	uid, err := mysql.SelectUserByUsername(commentUsername)
	fmt.Println(uid)
	if err != nil {
		fmt.Println("PostComment() service.article.SelectUserByUsername err=", err)
		return err
	}

	//判断评论类型
	switch commentType {
	//给文章评论
	case "article":
		//获取文章id
		aid, _ := readMessage.StringToInt(m["id"].(string))
		fmt.Println("aid:", aid)
		//向数据库插入评论数据
		err = mysql.InsertIntoCommentsForArticle(commentContent, aid, uid)

		if err != nil {
			return err
		}
	case "comment":
		//获取评论id
		//pid, err := utils.StringToInt(m["id"])
		pid := m["id"].(int)
		fmt.Println("pid:", pid)
		//向数据库插入评论数据
		err = mysql.InsertIntoCommentsForComment(commentContent, pid, int(uid))
		if err != nil {
			return err
		}
	}
	return nil
}
