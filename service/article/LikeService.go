package article

import (
	"fmt"
	"studentGrow/dao/redis"
)

// Like 点赞
func Like(objId, userId string, likeType int) error {
	err := redis.AddUserToLikeSet(redis.List[likeType]+objId, userId, likeType)
	if err != nil {
		fmt.Println("Like() service.article.likeService.AddUserToLikeSet err=", err)
		return err
	}
	likes, err := redis.GetObjLikes(objId, likeType)
	if err != nil {
		fmt.Println("Like() service.article.likeService.GetObjLikes err=", err)
		return err
	}
	if likes >= 0 {
		err = redis.SetObjLikes(objId, likes+1, likeType)
		if err != nil {
			fmt.Println("Like() service.article.likeService.SetObjLikes err=", err)
			return err
		}
	}

	return nil
}

// CancelLike 取消点赞
func CancelLike(objId, userId string, likeType int) error {

	ok, err := redis.IsUserLiked(objId, userId, likeType)
	if err != nil {
		fmt.Println("CancelLike() service.article.likeService.IsUserLiked err=", err)
		return err
	}
	if ok {
		err = redis.RemoveUserFromLikeSet(objId, userId, likeType)
		if err != nil {
			fmt.Println("CancelLike() service.article.likeService.RemoveUserFromLikeSet err=", err)
			return err
		}
		likes, err := redis.GetObjLikes(objId, likeType)
		if err != nil {
			fmt.Println("CancelLike() service.article.likeService.GetObjLikes err=", err)
			return err
		}
		if likes > 0 {
			err = redis.SetObjLikes(objId, likes-1, likeType)
			if err != nil {
				fmt.Println("CancelLike() service.article.likeService.SetObjLikes err=", err)
				return err
			}
		}
	}
	return nil
}

// LikeObjOrNot 检查是否点赞并点赞
func LikeObjOrNot(objId, userId string, likeType int) error {
	//获取当前点赞文章列表
	slice, err := redis.GetObjLikedUsers(objId, likeType)
	if err != nil {
		fmt.Println("LikeObjOrNot() service.article.likeService.GetObjLikedUsers err=", err)
		return err
	}
	likeUsers := make(map[string]struct{})
	for _, s := range slice {
		likeUsers[s] = struct{}{}
	}
	//若存在该用户，则取消点赞
	_, ok := likeUsers[userId]
	if len(likeUsers) > 0 && ok {
		err = CancelLike(objId, userId, likeType)
		if err != nil {
			fmt.Println("LikeObjOrNot() service.article.likeService.CancelLike err=", err)
			return err
		}
	} else {
		//反之，点赞
		err = Like(objId, userId, likeType)
		if err != nil {
			fmt.Println("LikeObjOrNot() service.article.likeService.Like err=", err)
			return err
		}
	}
	return nil
}
