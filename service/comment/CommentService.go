package comment

import (
	"fmt"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/dao/redis"
	"studentGrow/models/gorm_model"
)

// PostComment 发布评论
func PostComment(commentType, username, content string, id int) error {
	//类型comment_type:‘article’or‘comment’;id;comment_content;comment_username

	//获取用户id
	uid, err := mysql.SelectUserByUsername(username)
	fmt.Println(uid)
	if err != nil {
		fmt.Println("PostComment() service.article.SelectUserByUsername err=", err)
		return err
	}

	//判断评论类型
	switch commentType {
	//给文章评论
	case "article":
		//向数据库插入评论数据
		err = mysql.InsertIntoCommentsForArticle(content, id, uid)
		if err != nil {
			return err
		}
	case "comment":
		//向数据库插入评论数据
		err = mysql.InsertIntoCommentsForComment(content, id, uid)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetLel1CommentListService 获取一级评论列表
func GetLel1CommentListService(username string, aid int) ([]gorm_model.Comment, error) {
	// 获取文章对应的评论
	comments, err := mysql.QueryLevelOneComments(aid)
	if err != nil {
		fmt.Println("GetLel1CommentList() service.article.QueryLevelOneComments err=", err)
		return nil, err
	}

	// 该用户是否点赞
	for i := 0; i < len(comments); i++ {
		liked, err := redis.IsUserLiked(strconv.Itoa(int(comments[i].ID)), username, 1)
		if err != nil {
			fmt.Println("GetLel1CommentList() service.article.IsUserLiked err=", err)
			return nil, err
		}
		comments[i].IsLike = liked
	}

	return comments, nil
}

// GetLel2CommentListService 获取二级评论列表
func GetLel2CommentListService(username string, cid int) ([]gorm_model.Comment, error) {
	// 获取文章对应的评论
	comments, err := mysql.QueryLevelTwoComments(cid)
	if err != nil {
		fmt.Println("GetLel1CommentList() service.article.QueryLevelTwoComments err=", err)
		return nil, err
	}

	// 该用户是否点赞
	for i := 0; i < len(comments); i++ {
		liked, err := redis.IsUserLiked(strconv.Itoa(int(comments[i].ID)), username, 1)
		if err != nil {
			fmt.Println("GetLel1CommentList() service.article.IsUserLiked err=", err)
			return nil, err
		}
		comments[i].IsLike = liked
	}

	return comments, nil
}
