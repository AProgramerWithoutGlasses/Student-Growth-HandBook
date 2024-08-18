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
