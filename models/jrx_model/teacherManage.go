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
