package redis

import (
	"fmt"
	"strconv"
)

var Selection = "select"

// AddArticleToCollectSet 添加文章到用户收藏集合中
func AddArticleToCollectSet(uid string, aid string) error {
	if err := RDB.SAdd(Selection+uid, aid).Err(); err != nil {
		fmt.Println("AddArticleToSelectSet() service.article.SAdd err=", err)
		return err
	}
	return nil
}

// IsUserCollected 检查用户是否已经收藏
func IsUserCollected(uid string, aid string) (bool, error) {
	result, err := RDB.SIsMember(Selection+uid, aid).Result()
	if err != nil {
		fmt.Println("IsUserSelected() service.article.SIsMember err=", err)
		return false, nil
	}

	return result, nil
}

// SetArticleCollections 设置文章收藏数
func SetArticleCollections(aid string, selectNum int) error {
	if err := RDB.HIncrBy(Selection, aid, int64(selectNum)).Err(); err != nil {
		fmt.Println("IsUserSelected() service.article.SIsMember err=", err)
		return err
	}
	return nil
}

// GetArticleCollections 获取文章收藏数
func GetArticleCollections(aid string) (int, error) {
	selectNum, err := RDB.HGet(Selection, aid).Result()
	if err != nil {
		fmt.Println("GetArticleSelections() service.article.HGet err=", err)
		return -1, err
	}
	res, err := strconv.Atoi(selectNum)
	if err != nil {
		fmt.Println("GetArticleSelections() service.article.Atoi err=", err)
		return -1, err
	}
	return res, nil
}

// GetUserCollectionSet 获取用户的收藏集合
func GetUserCollectionSet(uid string) ([]string, error) {
	slice, err := RDB.SMembers(Selection + uid).Result()
	if err != nil {
		fmt.Println("GetUserSelectionSet() service.article.SMembers err=", err)
		return nil, err
	}

	return slice, nil
}

// RemoveUserCollectionSet 将文章从用户收藏集合中移除
func RemoveUserCollectionSet(aid, uid string) error {
	err := RDB.SRem(Selection+uid, aid).Err()
	if err != nil {
		fmt.Println("RemoveUserFromLikeSet() service.article.ArticleLikeService err=", err)
		return err
	}
	return nil
}
