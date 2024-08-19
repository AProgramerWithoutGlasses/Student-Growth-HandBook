package jrx_model

// 个人主页信息
type HomepageMesStruct struct {
	Username     string `json:"username"`
	Ban          bool   `json:"ban"`
	Name         string `json:"name"`
	UserHeadShot string `json:"user_headshot"`
	UserMotto    string `json:"user_motto"`
	UserFans     int    `json:"user_fans"`
	UserConcern  int    `json:"user_concern"`
	UserLike     int    `json:"user_like"`
	Point        int    `json:"user_point"`
	UserClass    string `json:"user_class"`
	UserIdentity string `json:"user_identity"`
}

// 个人资料信息
type HomepageDataStruct struct {
	Name         string `json:"name"`
	UserHeadShot string `json:"user_headshot"`
	UserMotto    string `json:"user_motto"`
	UserClass    string `json:"user_class"`
	UserGender   string `json:"user_gender"`
	Phone_number string `json:"phone_number"`
	UserEmail    string `json:"user_email"`
	UserYear     string `json:"user_year"`
}
