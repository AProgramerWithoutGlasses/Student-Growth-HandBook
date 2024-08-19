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

// 用于粉丝列表
type HomepageFanStruct struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Motto    string `json:"user_motto"`
	HeadShot string `json:"user_headshot"`
}

type HomepageArticleHistoryStruct struct {
	ID            string `json:"article_id"`
	HeadShot      string `json:"user_headshot"`
	Name          string `json:"name"`
	Content       string `json:"article_content"`
	Pic           string `json:"article_pic"`
	CommentAmount int    `json:"comment_amount"`
	LikeAmount    int    `json:"like_amount"`
}
