package jrx_model

// 学生信息表（为贴合apifox的字段，备用）
type StuMesStruct struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Class       string `json:"class"`
	Year        string `json:"year"`
	Gender      string `json:"gender"`
	Telephone   string `json:"telephone"`
	Ban         bool   `json:"ban"`
	ManagerType string `json:"manager_type"`
}

// 用于入学年份下拉框
type YearStruct struct {
	Id_Year string `json:"value"`
	Year    string `json:"label"`
}

// 用于班级下拉框
type ClassStruct struct {
	Id_class string `json:"value"`
	Class    string `json:"label"`
}

// ResponseStruct 返回查询结果给前端
type ResponseStruct struct {
	Year            []YearStruct   `json:"year"`
	Class           []ClassStruct  `json:"class"`
	StuInfo         []StuMesStruct `json:"stuInfo"`
	AllStudentCount int            `json:"allStudentCount"`
}

// queryParmaStruct 用于获取查询参数
type QueryParmaStruct struct {
	Year          string `json:"year"`
	Class         string `json:"class"`
	Gender        string `json:"gender"`
	IsDisable     bool   `json:"is_disable"`
	SearchSelect  string `json:"search_select"`
	SearchMessage string `json:"search_message"`
}

// 修改学生信息
type ChangeStuMesStruct struct {
	Username     string `json:"username" form:"username"`
	Class        string `json:"class" form:"class"`
	Phone_number string `json:"telephone" form:"phone_number"`
	Password     string `json:"password" form:"password"`
}

// 学生信息表（为贴合apifox的字段，备用）(year int)
type StuMesYearIntStruct struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Class     string `json:"class"`
	Year      int    `json:"year"`
	Gender    string `json:"gender"`
	Telephone string `json:"telephone"`
	Ban       bool   `json:"ban"`
	IsManager bool   `json:"isManager"`
}

type SelectedStuMesStruct struct {
	Selected_students []StuMesStruct `json:"selected_students"`
}
