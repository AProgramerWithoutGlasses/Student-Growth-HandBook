package service

import (
	"fmt"
	"studentGrow/dao/mysql"
	"studentGrow/models/jrx_model"
)

func GetHomepageMesService(username string) (*jrx_model.HomepageMesStruct, error) {
	homepageMes := &jrx_model.HomepageMesStruct{}

	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return nil, err
	}

	// 从user表中获取数据
	userMes, err := mysql.GetHomepageUserMesDao(id)
	if err != nil {
		return nil, err
	}

	// 从其他表中获取数据
	userFans, err := mysql.GetHomepageFansCountDao(id)
	if err != nil {
		return nil, err
	}

	userConcern, err := mysql.GetHomepageConcernCountDao(id)
	if err != nil {
		return nil, err
	}

	userLike, err := mysql.GetHomepageLikeCountDao(id)
	if err != nil {
		return nil, err
	}

	// 将获得的数据存储到 homepageMesStruct中
	homepageMes.Username = userMes.Username
	homepageMes.Ban = userMes.Ban
	homepageMes.Name = userMes.Name
	homepageMes.UserHeadShot = userMes.HeadShot
	homepageMes.UserMotto = userMes.Motto
	homepageMes.UserFans = userFans
	homepageMes.UserConcern = userConcern
	homepageMes.UserLike = userLike
	homepageMes.Point = userMes.Point
	homepageMes.UserClass = userMes.Class
	homepageMes.UserIdentity = userMes.Identity

	fmt.Printf("homepageMes2: %+v\n", homepageMes)

	return homepageMes, err
}
