package service

import (
	"fmt"
	"mime/multipart"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/models/jrx_model"
	"studentGrow/utils/fileProcess"
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

	point, err := mysql.GetHomepagePointDao(id)
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
	homepageMes.Point = point
	homepageMes.UserClass = userMes.Class

	fmt.Printf("homepageMes2: %+v\n", homepageMes)

	return homepageMes, err
}

func UpdateHomepageMottoService(username string, motto string) error {
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return err
	}

	err = mysql.UpdateHomepageMottoDao(id, motto)
	if err != nil {
		return err
	}

	return nil
}

func UpdateHomepagePhoneNumberService(username string, phone_number string) error {
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return err
	}

	err = mysql.UpdateHomepagePhoneNumberDao(id, phone_number)
	if err != nil {
		return err
	}

	return nil
}

func UpdateHomepageEmailService(username string, email string) error {
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return err
	}

	err = mysql.UpdateHomepageEmailDao(id, email)
	if err != nil {
		return err
	}

	return nil
}

func GetHomepageUserDataService(username string) (*jrx_model.HomepageDataStruct, error) {
	userData := &jrx_model.HomepageDataStruct{}
	userDataTemp := &gorm_model.User{}

	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return nil, err
	}

	userDataTemp, err = mysql.GetHomepageUserMesDao(id)
	if err != nil {
		return nil, err
	}

	userData.Name = userDataTemp.Name
	userData.UserHeadShot = userDataTemp.HeadShot
	userData.UserGender = userDataTemp.Gender
	userData.UserClass = userDataTemp.Class
	userData.UserMotto = userDataTemp.Motto
	userData.Phone_number = userDataTemp.PhoneNumber
	userData.UserEmail = userDataTemp.MailBox
	userData.UserYear = userDataTemp.PlusTime.Format("2006")

	fmt.Printf("userData : %+v\n", userData)

	return userData, err
}

// 更新个人主页头像
func UpdateHeadshotService(file *multipart.FileHeader, username string) error {
	url, err := fileProcess.UploadFile("image", file)
	if err != nil {
		return err
	}

	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return err
	}

	err = mysql.UpdateHeadshotDao(id, url)
	if err != nil {
		return err
	}

	return err
}

func GetFansListService(username string) ([]jrx_model.HomepageFanStruct, error) {
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return nil, err
	}

	fansIdList, err := mysql.GetFansIdListDao(id)
	if err != nil {
		return nil, err
	}

	fansList, err := mysql.GetFansListDao(fansIdList)
	if err != nil {
		return nil, err
	}

	return fansList, err
}

func GetConcernListService(username string) ([]jrx_model.HomepageFanStruct, error) {
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return nil, err
	}

	concernIdList, err := mysql.GetConcernIdListDao(id)
	if err != nil {
		return nil, err
	}

	concernList, err := mysql.GetConcernListDao(concernIdList)
	if err != nil {
		return nil, err
	}

	return concernList, err
}

func ChangeConcernService(username string, othername string) error {
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return err
	}

	otherId, err := mysql.GetIdByUsername(othername)
	if err != nil {
		return err
	}

	err = mysql.ChangeConcernDao(id, otherId)
	if err != nil {
		return err
	}

	return err
}

func GetHistoryService(page int, limit int, username string) ([]jrx_model.HomepageArticleHistoryStruct, error) {
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return nil, err
	}

	homepageArticleHistoryList, err := mysql.GetHistoryByArticleDao(id, page, limit)
	if err != nil {
		return nil, err
	}

	return homepageArticleHistoryList, err
}

func GetStarService(page int, limit int, username string) ([]jrx_model.HomepageArticleHistoryStruct, error) {
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return nil, err
	}

	homepageStarList, err := mysql.GetStarDao(id, page, limit)
	if err != nil {
		return nil, err
	}

	return homepageStarList, err
}

func GetClassListService(username string) ([]jrx_model.HomepageClassmateStruct, error) {
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return nil, err
	}

	class, err := mysql.GetClassById(id)
	if err != nil {
		return nil, err
	}

	classmateList, err := mysql.GetClassmateList(class)
	if err != nil {
		return nil, err
	}

	return classmateList, err
}
