package routesJoinAudit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nguyenthenguyen/docx"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
)

type resList struct {
	ListName string                   `json:"list_name"`
	IsFinish bool                     `json:"is_finish"`
	List     []map[string]interface{} `json:"list"`
}

type rec struct {
	ActivityID int    `json:"activity_id"`
	CurMenu    string `json:"cur_menu"`
	Title      string `json:"title"`
}

func ExportFormMsg(c *gin.Context) {
	token := token2.NewToken(c)
	_, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	var cr rec
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "json解析失败")
		return
	}
	userMsgList, classList := mysql.UserListWithOrganizer(cr.ActivityID, cr.CurMenu)
	var isFinish = true
	if len(userMsgList) == 0 {
		isFinish = false
	}
	if cr.Title == "" {
		cr.Title = "通过名单"
	}
	rewriteTemplate(userMsgList)
	filePath := UseWordTemplateMakeDocx(cr.Title, userMsgList, classList)
	response.ResponseSuccess(c, resList{
		ListName: cr.CurMenu,
		IsFinish: isFinish,
	})
	c.FileAttachment(filePath, cr.Title+".docx")
	fmt.Println(filePath)
}
func UseWordTemplateMakeDocx(title string, List []map[string]interface{}, classList []string) (filePath string) {
	classified, replaceNameList := rewriteTemplate(List)
	fmt.Println(classified)
	fmt.Println(replaceNameList)
	r, err := docx.ReadDocxFile("controller/routesJoinAudit/word/rewriteTemplate.docx")
	if err != nil {
		panic(err)
	}
	defer r.Close()
	docx1 := r.Editable()
	var classNameList []string
	for k, v := range classList {
		classStr := "CLASS"
		nameStr := "NAME"
		classStr = "{{" + classStr + strconv.Itoa(k) + "}}"
		nameStr = "{{" + nameStr + strconv.Itoa(k) + "}}"
		err = docx1.Replace(classStr, v+":", -1)
		if err != nil {
			panic(err)
		}
		classNameList = make([]string, 0)
		for _, value := range classified[v] {
			userName, _ := value["name"].(string)
			classNameList = append(classNameList, userName)
		}
		rewriteList := strings.Join(classNameList, "\t")
		err = docx1.Replace(nameStr, rewriteList, -1)
		if err != nil {
			panic(err)
		}
	}
	err = docx1.Replace("{{Title}}", title, -1)
	if err != nil {
		panic(err)
	}
	filePath = "controller/routesJoinAudit/word/" + fmt.Sprintf("%s", title) + ".docx"
	err = docx1.WriteToFile(filePath)
	if err != nil {
		zap.L().Error(err.Error())
	}
	return
}

func rewriteTemplate(dynamicList []map[string]interface{}) (classified map[string][]map[string]interface{}, replaceNameList []string) {
	classified = make(map[string][]map[string]interface{})
	for _, v := range dynamicList {
		userClass, _ := v["user_class"].(string)
		classified[userClass] = append(classified[userClass], v)
	}

	n := len(classified)
	replaceNameList = make([]string, 0)
	for i := 0; i < n; i++ {
		classStr := "CLASS"
		nameStr := "NAME"
		classStr = "{{" + classStr + strconv.Itoa(i) + "}}"
		nameStr = "{{" + nameStr + strconv.Itoa(i) + "}}"
		replaceNameList = append(replaceNameList, classStr)
		replaceNameList = append(replaceNameList, nameStr)
	}
	rewriteList := strings.Join(replaceNameList, "\n")
	// 读取模板文件
	r, err := docx.ReadDocxFile("controller/routesJoinAudit/word/template.docx")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	// 使文档可编辑
	docx1 := r.Editable()

	// 替换模板中的占位符
	err = docx1.Replace("{{LIST}}", rewriteList, -1)
	if err != nil {
		panic(err)
	}
	// 将修改后的文档写入新文件
	err = docx1.WriteToFile("controller/routesJoinAudit/word/rewriteTemplate.docx")
	if err != nil {
		panic(err)
	}
	return
}
