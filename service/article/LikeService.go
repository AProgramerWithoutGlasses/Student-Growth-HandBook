package article

import (
	"fmt"
	"strconv"
	"studentGrow/dao/redis"
)

var list = []string{"article_like-", "comment_like-"}

// AddUserToLikeSet 添加用户到文章或评论点赞集合中
func AddUserToLikeSet(objId, userId string, likeType int) {

	err := redis.RDB.SAdd(list[likeType]+objId, userId).Err()

	if err != nil {
		fmt.Println("AddUserToLikeSet() service.article.ArticleLikeService err=", err)
		return
	}
}

// IsUserLiked 检查用户是否已经点赞
func IsUserLiked(objId, userId string, likeType int) bool {
	res, err := redis.RDB.SIsMember(list[likeType]+objId, userId).Result()
	if err != nil {
		fmt.Println("IsUserLiked() service.article.ArticleLikeService err=", err)
		return false
	}
	return res
}

// SetObjLikes 设置文章或评论的点赞数
func SetObjLikes(objId string, likeNum int, likeType int) {
	err := redis.RDB.HIncrBy(list[likeType], objId, int64(likeNum)).Err()
	if err != nil {
		fmt.Println("SetArticleLikes() service.article.ArticleLikeService err=", err)
		return
	}
}

// GetObjLikes 获取文章或评论点赞数
func GetObjLikes(objId string, likeType int) int {
	likesNumResult, err := redis.RDB.HGet(list[likeType], objId).Result()
	if err != nil {
		fmt.Println("GetArticleLikes() service.article.ArticleLikeService err=", err)
		return 0
	}
	res, err := strconv.Atoi(likesNumResult)
	if err != nil {
		fmt.Println("GetArticleLikes() service.article.ArticleLikeService err=", err)
		return 0
	}
	return res
}

// GetObjLikedUsers 获取文章或评论点赞的用户username集合
func GetObjLikedUsers(objId string, likeType int) (result []string) {
	slice, err := redis.RDB.SMembers(list[likeType] + objId).Result()

	if err != nil {
		fmt.Println("GetArticleLikedUsers() service.article.ArticleLikeService err=", err)
		return nil
	}
	return slice
}

// RemoveUserFromLikeSet 移除用户从文章或评论的点赞集合中
func RemoveUserFromLikeSet(objId, userId string, likeType int) {
	err := redis.RDB.SRem(list[likeType]+objId, userId).Err()
	if err != nil {
		fmt.Println("RemoveUserFromLikeSet() service.article.ArticleLikeService err=", err)
		return
	}
}

// Like 点赞
func Like(objId, userId string, likeType int) {
	AddUserToLikeSet(list[likeType]+objId, userId, likeType)
	likes := GetObjLikes(objId, likeType)
	if likes >= 0 {
		SetObjLikes(objId, likes+1, likeType)
	}
}

// CancelLike 取消点赞
func CancelLike(objId, userId string, likeType int) {
	if IsUserLiked(objId, userId, likeType) {
		RemoveUserFromLikeSet(objId, userId, likeType)
		likes := GetObjLikes(objId, likeType)
		if likes > 0 {
			SetObjLikes(objId, likes-1, likeType)
		}
	}
}

// LikeObjOrNot 检查是否点赞
func LikeObjOrNot(objId, userId string, likeType int) bool {
	//获取当前点赞文章列表
	slice := GetObjLikedUsers(objId, likeType)
	likeUsers := make(map[string]struct{})
	for _, s := range slice {
		likeUsers[s] = struct{}{}
	}
	//若存在该用户，则取消点赞
	_, ok := likeUsers[userId]
	if len(likeUsers) > 0 && ok {
		CancelLike(objId, userId, likeType)
		return false
	} else {
		//反之，点赞
		Like(objId, userId, likeType)
		return true
	}
}
