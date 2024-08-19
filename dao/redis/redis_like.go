package redis

import (
	"fmt"
	redis2 "github.com/go-redis/redis"
	"strconv"
)

var List = []string{"article", "comment"}

// AddUserToLikeSet 添加用户到文章或评论点赞集合中
func AddUserToLikeSet(objId, userId string, likeType int, pipe redis2.Pipeliner) error {

	err := pipe.SAdd(List[likeType]+objId, userId).Err()

	if err != nil {
		fmt.Println("AddUserToLikeSet() service.article.ArticleLikeService err=", err)
		return err
	}

	return nil
}

// IsUserLiked 检查用户是否已经点赞
func IsUserLiked(objId, userId string, likeType int) (bool, error) {
	res, err := RDB.SIsMember(List[likeType]+objId, userId).Result()
	if err != nil {
		fmt.Println("IsUserLiked() service.article.ArticleLikeService err=", err)
		return false, err
	}
	return res, err
}

// SetObjLikes 设置文章或评论的点赞数
func SetObjLikes(objId string, likeNum int, likeType int, pipe redis2.Pipeliner) error {
	err := pipe.HIncrBy(List[likeType], objId, int64(likeNum)).Err()
	if err != nil {
		fmt.Println("SetArticleLikes() service.article.ArticleLikeService err=", err)
		return err
	}
	return nil
}

// GetObjLikes 获取文章或评论点赞数
func GetObjLikes(objId string, likeType int, pipe redis2.Pipeliner) (int, error) {
	likesNumResult, err := pipe.HGet(List[likeType], objId).Result()
	if err != nil {
		fmt.Println("GetArticleLikes() service.article.ArticleLikeService err=", err)
		return -1, err
	}
	res, err := strconv.Atoi(likesNumResult)
	if err != nil {
		fmt.Println("GetArticleLikes() service.article.ArticleLikeService err=", err)
		return -1, err
	}
	return res, nil
}

// GetObjLikedUsers 获取文章或评论点赞的用户username集合
func GetObjLikedUsers(objId string, likeType int) (result []string, err error) {
	slice, err := RDB.SMembers(List[likeType] + objId).Result()

	if err != nil {
		fmt.Println("GetArticleLikedUsers() service.article.ArticleLikeService err=", err)
		return nil, err
	}
	return slice, nil
}

// RemoveUserFromLikeSet 移除用户从文章或评论的点赞集合中
func RemoveUserFromLikeSet(objId, userId string, likeType int) error {
	err := RDB.SRem(List[likeType]+objId, userId).Err()
	if err != nil {
		fmt.Println("RemoveUserFromLikeSet() service.article.ArticleLikeService err=", err)
		return err
	}

	return nil
}

//// QueryCommentLikeNum 查询评论点赞数
//func QueryCommentLikeNum(cid int) (int, error) {
//	num, err := GetObjLikes(strconv.Itoa(cid), 1)
//	if err != nil {
//		fmt.Println("QueryCommentLikeNum() dao.mysql.nzx_sql.GetObjLikes err=", err)
//		return 0, err
//	}
//	return num, nil
//}
//
//// UpdateCommentLikeNum 设置评论点赞数
//func UpdateCommentLikeNum(cid, num int) error {
//	err := SetObjLikes(strconv.Itoa(cid), num, 1)
//	if err != nil {
//		fmt.Println("UpdateCommentLikeNum() dao.mysql.nzx_sql.SetObjLikes err=", err)
//		return err
//	}
//	return nil
//}
