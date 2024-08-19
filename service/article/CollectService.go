package article

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/dao/redis"
	"studentGrow/models/nzx_model"
)

// CollectService Collect 收藏
func CollectService(username, aid string) error {
	// 收藏文章
	err := redis.AddArticleToCollectSet(username, aid)
	if err != nil {
		fmt.Println("Collect() service.article.AddArticleToCollectSet err=", err)
		return err
	}
	// 获取文章收藏数
	selections, err := redis.GetArticleCollections(aid)
	if err != nil {
		fmt.Println("Collect() service.article.GetArticleCollections err=", err)
		return err
	}
	// 收藏数+1
	if selections >= 0 {
		err = redis.SetArticleCollections(aid, selections+1)
		if err != nil {
			fmt.Println("Collect() service.article.SetArticleCollections err=", err)
			return err
		}
	}

	// 写入通道
	articleId, err := strconv.Atoi(aid)
	if err != nil {
		fmt.Println("Collect() service.article.Atoi err=", err)
		return err
	}
	ArticleCollectChan <- nzx_model.RedisCollectData{Aid: articleId, Username: username, Operator: "collect"}
	return nil
}

// CancelCollectService CancelCollect 取消收藏
func CancelCollectService(aid, username string) error {
	isExist, err := redis.IsUserCollected(username, aid)
	if err != nil {
		fmt.Println("CancelCollect() service.article.IsUserCollected err=", err)
		return err
	}

	if isExist {
		err = redis.RemoveUserCollectionSet(aid, username)
		if err != nil {
			fmt.Println("CancelCollect() service.article.RemoveUserCollectionSet err=", err)
			return err
		}
		selections, err := redis.GetArticleCollections(aid)
		if err != nil {
			fmt.Println("CancelCollect() service.article.RemoveUserCollectionSet err=", err)
			return err
		}
		if selections > 0 {
			err := redis.SetArticleCollections(aid, selections-1)
			if err != nil {
				fmt.Println("CancelCollect() service.article.SetArticleCollections err=", err)
				return err
			}
		}

		// 收藏数-1
		// 获取文章收藏数
		selections, err = redis.GetArticleCollections(aid)
		if err != nil {
			fmt.Println("Collect() service.article.GetArticleCollections err=", err)
			return err
		}
		// 收藏数-1
		if selections >= 0 {
			err = redis.SetArticleCollections(aid, selections-1)
			if err != nil {
				fmt.Println("Collect() service.article.SetArticleCollections err=", err)
				return err
			}
		}

		// 写入通道
		articleId, err := strconv.Atoi(aid)
		if err != nil {
			fmt.Println("Collect() service.article.Atoi err=", err)
			return err
		}
		ArticleCollectChan <- nzx_model.RedisCollectData{Aid: articleId, Username: username, Operator: "cancel_collect"}
	}

	return nil
}

// CollectOrNotService CollectOrNot 检查是否收藏并收藏或取消收藏
func CollectOrNotService(aid, username string) error {
	// 获取当前用户收藏列表
	slice, err := redis.GetUserCollectionSet(username)
	if err != nil {
		fmt.Println("CollectOrNot() service.article.GetUserCollectionSet err=", err)
		return err
	}

	// 若存在该文章,则取消收藏
	selectArticles := make(map[string]struct{})
	for _, s := range slice {
		selectArticles[s] = struct{}{}
	}
	_, ok := selectArticles[aid]

	if len(selectArticles) > 0 && ok {
		err = CancelCollectService(aid, username)
		if err != nil {
			fmt.Println("CollectOrNot() service.article.GetUserCollectionSet err=", err)
			return err
		}
	} else {
		// 反之，收藏
		err = CollectService(username, aid)
		if err != nil {
			fmt.Println("CollectOrNot() service.article.GetUserCollectionSet err=", err)
			return err
		}
	}
	return nil
}

/*
mysql
*/

// CollectToMysql 收藏
func CollectToMysql(aid int, username string) error {
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		fmt.Println("CollectToMysql() service.article.GetIdByUsername err=", err)
		return err
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 添加收藏记录
		err = mysql.InsertCollectRecord(aid, uid, tx)
		if err != nil {
			fmt.Println("CollectToMysql() service.article.InsertCollectRecord err=", err)
			return err
		}

		// 获取收藏数
		num, err := mysql.QueryCollectNum(aid)
		if err != nil {
			fmt.Println("CollectToMysql() service.article.QueryCollectNum err=", err)
			return err
		}
		// 收藏数+1
		err = mysql.UpdateCollectNum(aid, num+1, tx)
		if err != nil {
			fmt.Println("CollectToMysql() service.article.UpdateCollectNum err=", err)
			return err
		}
		return nil
	})
	if err != nil {
		zap.L().Error("CollectToMysql() service.article.Transaction err=", zap.Error(err))
		return err
	}
	return nil
}

// CancelCollectToMysql 取消收藏
func CancelCollectToMysql(aid int, username string) error {
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		fmt.Println("CollectToMysql() service.article.GetIdByUsername err=", err)
		return err
	}
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 删除收藏记录
		err = mysql.DeleteCollectRecord(aid, uid, tx)
		if err != nil {
			fmt.Println("CollectToMysql() service.article.InsertCollectRecord err=", err)
			return err
		}

		// 获取收藏数
		num, err := mysql.QueryCollectNum(aid)
		if err != nil {
			fmt.Println("CollectToMysql() service.article.QueryCollectNum err=", err)
			return err
		}
		// 收藏数-1
		err = mysql.UpdateCollectNum(aid, num-1, tx)
		if err != nil {
			fmt.Println("CollectToMysql() service.article.UpdateCollectNum err=", err)
			return err
		}
		return nil
	})
	if err != nil {
		zap.L().Error("CancelCollectToMysql() service.article.Transaction err=", zap.Error(err))
		return err
	}
	return nil
}
