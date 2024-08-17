package article

import (
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
)

func UpdatePointService(uid, point, topicId int) error {
	// 获取当前分数
	curPoint, err := mysql.QueryUserPointByTopic(topicId, uid)
	if err != nil {
		zap.L().Error("PublishArticleService() service.article.QueryUserPointByTopic err=", zap.Error(err))
		return err
	}
	// 修改分数
	err = mysql.UpdateUserPointByTopic(curPoint+point, uid, topicId)
	if err != nil {
		zap.L().Error("PublishArticleService() service.article.UpdateUserPointByTopic err=", zap.Error(err))
		return err
	}
	return nil
}

// UpdatePointByUsernamePointAid 已知username,point,aid修改分数
func UpdatePointByUsernamePointAid(username string, point, aid int) error {
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("UpdatePointByUsernamePointAid() service.article.GetIdByUsername err=", zap.Error(err))
		return err
	}
	//  获取topicId
	topic, err := mysql.QueryTopicByArticleId(aid)
	if err != nil {
		zap.L().Error("UpdatePointByUsernamePointAid() service.article.QueryTopicByArticleId err=", zap.Error(err))
		return err
	}
	topicId, err := mysql.QueryTopicIdByTopicName(topic)
	if err != nil {
		zap.L().Error("UpdatePointByUsernamePointAid() service.article.QueryTopicIdByTopicName err=", zap.Error(err))
		return err
	}

	err = UpdatePointService(uid, point, topicId)
	if err != nil {
		zap.L().Error("UpdatePointByUsernamePointAid() service.article.UpdatePointService err=", zap.Error(err))
		return err
	}
	return nil
}
