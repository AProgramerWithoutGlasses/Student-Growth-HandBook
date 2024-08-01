package comment

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"studentGrow/dao/mysql"
)

// PostComment 发布评论
func PostComment(j *jsonvalue.V) error {
	//类型comment_type:‘article’or‘comment’;id;comment_content;comment_username

	//获取数据
	commentType, _ := j.GetString("comment_type")
	commentContent, _ := j.GetString("comment_content")
	commentUsername, _ := j.GetString("comment_username")

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
		aid, _ := j.GetInt("id")
		fmt.Println("aid:", aid)
		//向数据库插入评论数据
		err = mysql.InsertIntoCommentsForArticle(commentContent, aid, uid)

		if err != nil {
			return err
		}
	case "comment":
		//获取评论id
		//pid, err := utils.StringToInt(m["id"])
		pid, _ := j.GetInt("id")
		fmt.Println("pid:", pid)
		//向数据库插入评论数据
		err = mysql.InsertIntoCommentsForComment(commentContent, pid, uid)
		if err != nil {
			return err
		}
	}
	return nil
}
