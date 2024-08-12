package mysql

// InsertIntoCommentsForArticle 向数据库插入评论数据(回复文章)

func InsertIntoCommentsForArticle(content string, aid int, uid int) (err error) {
	//content;id;username
	//comment := model.Comment{
	//	Model:        gorm.Model{},
	//	Content:      content,
	//	UpvoteAmount: 0,
	//	IsRead:       false,
	//	Del:          false,
	//	Uid:          uid,
	//	Pid:          0,
	//	Aid:          aid,
	//	Upvote:       nil,
	//}
	//DB.Create(&comment)

	return nil
}

// InsertIntoCommentsForComment 向数据库插入评论数据(回复评论)
func InsertIntoCommentsForComment(content string, uid int, pid int) (err error) {
	//content;id;username
	//comment := model.Comment{
	//	Model:        gorm.Model{},
	//	Content:      content,
	//	UpvoteAmount: 0,
	//	IsRead:       false,
	//	Del:          false,
	//	Uid:          uid,
	//	Pid:          pid,
	//	Aid:          0,
	//	Upvote:       nil,
	//}
	//
	//DB.Create(&comment)
	return nil
}
