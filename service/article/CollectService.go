package article

import (
	"fmt"
	"studentGrow/dao/redis"
)

// SelectService Select 收藏
func SelectService(uid, aid string) error {
	// 收藏文章
	err := redis.AddArticleToSelectSet(uid, aid)
	if err != nil {
		fmt.Println("Select() service.article.AddArticleToSelectSet err=", err)
		return err
	}
	// 获取文章收藏数
	selections, err := redis.GetArticleSelections(aid)
	if err != nil {
		fmt.Println("Select() service.article.GetArticleSelections err=", err)
		return err
	}
	// 收藏数+1
	if selections >= 0 {
		err = redis.SetArticleSelections(aid, selections+1)
		if err != nil {
			fmt.Println("Select() service.article.SetArticleSelections err=", err)
			return err
		}
	}
	return nil
}

// CancelSelectService CancelSelect 取消收藏
func CancelSelectService(aid, uid string) error {
	isExist, err := redis.IsUserSelected(uid, aid)
	if err != nil {
		fmt.Println("CancelSelect() service.article.IsUserSelected err=", err)
		return err
	}

	if isExist {
		err = redis.RemoveUserSelectionSet(aid, uid)
		if err != nil {
			fmt.Println("CancelSelect() service.article.RemoveUserSelectionSet err=", err)
			return err
		}
		selections, err := redis.GetArticleSelections(aid)
		if err != nil {
			fmt.Println("CancelSelect() service.article.RemoveUserSelectionSet err=", err)
			return err
		}
		if selections > 0 {
			err := redis.SetArticleSelections(aid, selections-1)
			if err != nil {
				fmt.Println("CancelSelect() service.article.SetArticleSelections err=", err)
				return err
			}
		}
	}

	return nil
}

// SelectOrNotService SelectOrNot 检查是否收藏并收藏或取消收藏
func SelectOrNotService(aid, uid string) error {
	// 获取当前用户收藏列表
	slice, err := redis.GetUserSelectionSet(uid)
	if err != nil {
		fmt.Println("SelectOrNot() service.article.GetUserSelectionSet err=", err)
		return err
	}

	selectArticles := make(map[string]struct{})
	for _, s := range slice {
		selectArticles[s] = struct{}{}
	}

	// 若存在该文章,则收藏
	_, ok := selectArticles[aid]
	if len(selectArticles) > 0 && ok {
		err = CancelSelectService(aid, uid)
		if err != nil {
			fmt.Println("SelectOrNot() service.article.GetUserSelectionSet err=", err)
			return err
		}
	} else {
		// 反之，收藏
		err = SelectService(uid, aid)
		if err != nil {
			fmt.Println("SelectOrNot() service.article.GetUserSelectionSet err=", err)
			return err
		}
	}
	return nil
}
