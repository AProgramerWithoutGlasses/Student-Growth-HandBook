package jrx_model

type QueryTeacherParamStruct struct {
	Gender        string `json:"user_gender"`
	Ban           *bool  `json:"user_ban"`
	IsManager     *bool  `json:"user_is_manager"`
	SearchSelect  string `json:"search_select"`
	SearchMessage string `json:"search_message"`
	Page          int    `json:"page"`
	Limit         int    `json:"limit"`
}

type QueryTeacherResStruct struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Gender    string `json:"user_gender"`
	Ban       *bool  `json:"user_ban"`
	IsManager *bool  `json:"user_is_manager"`
}

// 修改老师信息
type ChangeTeacherMesStruct struct {
	Name        string `json:"name" form:"name"`
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	Gender      string `json:"user_gender" form:"user_gender"`
	OldUsername string `json:"old_username" form:"old_username"`
}
