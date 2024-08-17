package jrx_model

type QueryTeacherParamStruct struct {
	Gender        string `json:"gender"`
	Ban           *bool  `json:"ban"`
	IsManager     *bool  `json:"isManager"`
	SearchSelect  string `json:"searchSelect"`
	SearchMessage string `json:"searchMessage"`
	Page          int    `json:"page"`
	Limit         int    `json:"limit"`
}

type QueryTeacherResStruct struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Ban       *bool  `json:"ban"`
	IsManager *bool  `json:"isManager"`
}
