package routesJoinAudit

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nguyenthenguyen/docx"
	"go.uber.org/zap"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
)

//go:embed template.docx
var data embed.FS

type resList struct {
	ListName string `json:"list_name"`
	IsFinish bool   `json:"is_finish"`
}

type rec struct {
	ActivityID int    `json:"activity_id"`
	CurMenu    string `json:"cur_menu"`
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
	userMsgList, _ := mysql.UserListWithOrganizer(cr.ActivityID, cr.CurMenu)
	var isFinish = true
	if len(userMsgList) == 0 {
		isFinish = false
	}
	response.ResponseSuccess(c, resList{
		ListName: cr.CurMenu,
		IsFinish: isFinish,
	})

}
func ExportFormFile(c *gin.Context) {
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
	_, _, activityMsg := mysql.OpenActivityStates()
	title := activityMsg.ActivityName
	if title == "" {
		title = "通过名单"
	}
	userMsgList, classList := mysql.UserListWithOrganizer(cr.ActivityID, cr.CurMenu)
	responseDOCX(userMsgList, c, activityMsg.ActivityName, classList)
	return
}

//func UseWordTemplateMakeDocx(title string, List []map[string]interface{}, classList []string) (filePath string) {
//	classified := rewriteTemplate(List)
//
//	r, err := docx.ReadDocxFile("controller/routesJoinAudit/word/rewriteTemplate.docx")
//	if err != nil {
//		zap.L().Error(err.Error())
//	}
//	defer r.Close()
//	docx1 := r.Editable()
//	var classNameList []string
//	for k, v := range classList {
//		classStr := "CLASS"
//		nameStr := "NAME"
//		classStr = "{{" + classStr + strconv.Itoa(k) + "}}"
//		nameStr = "{{" + nameStr + strconv.Itoa(k) + "}}"
//		err = docx1.Replace(classStr, v+":", -1)
//		if err != nil {
//			zap.L().Error(err.Error())
//		}
//		classNameList = make([]string, 0)
//		for _, value := range classified[v] {
//			userName, _ := value["name"].(string)
//			classNameList = append(classNameList, userName)
//		}
//		rewriteList := strings.Join(classNameList, "\t")
//		err = docx1.Replace(nameStr, rewriteList, -1)
//		if err != nil {
//			zap.L().Error(err.Error())
//		}
//	}
//	err = docx1.Replace("{{Title}}", title, -1)
//	if err != nil {
//		zap.L().Error(err.Error())
//	}
//	filePath = "controller/routesJoinAudit/word/" + fmt.Sprintf("%s", title) + ".docx"
//	err = docx1.WriteToFile(filePath)
//	if err != nil {
//		zap.L().Error(err.Error())
//	}
//	return
//}

func responseDOCX(dynamicList []map[string]interface{}, c *gin.Context, title string, classList []string) {
	classified := make(map[string][]map[string]interface{})
	for _, v := range dynamicList {
		userClass, _ := v["user_class"].(string)
		classified[userClass] = append(classified[userClass], v)
	}
	n := len(classified)
	replaceNameList := make([]string, 0)
	for i := 0; i < n; i++ {
		classStr := "CLASS"
		nameStr := "NAME"
		classStr = "{{" + classStr + strconv.Itoa(i) + "}}"
		nameStr = "{{" + nameStr + strconv.Itoa(i) + "}}"
		replaceNameList = append(replaceNameList, "\n"+classStr)
		replaceNameList = append(replaceNameList, nameStr)
	}
	rewriteList := strings.Join(replaceNameList, "\n")
	// 读取模板文件
	//读取文件夹dir
	templateFile, _ := data.Open("template.docx")
	defer templateFile.Close()
	// 创建新文件
	newFile, err := os.Create("./template.docx")
	if err != nil {
		zap.L().Error(err.Error())
	}
	defer newFile.Close()

	// 将嵌入文件的内容复制到新文件中
	_, err = io.Copy(newFile, templateFile)
	if err != nil {
		zap.L().Error(err.Error())
	}
	//entry为遍历dir文件夹中的每个文件

	r, err := docx.ReadDocxFile("./template.docx")
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, "文件打开失败")
		zap.L().Error(err.Error())
		return
	}
	defer r.Close()
	// 使文档可编辑
	docx1 := r.Editable()
	// 替换模板中的占位符
	err = docx1.Replace("{{LIST}}", rewriteList, -1)
	if err != nil {
		zap.L().Error(err.Error())
	}
	fmt.Println(rewriteList)
	// 将修改后的文档写入新文件
	var classNameList []string
	for k, v := range classList {
		classStr := "CLASS"
		nameStr := "NAME"
		classStr = "{{" + classStr + strconv.Itoa(k) + "}}"
		nameStr = "{{" + nameStr + strconv.Itoa(k) + "}}"
		err = docx1.Replace(classStr, v+":", -1)
		if err != nil {
			zap.L().Error(err.Error())
		}
		classNameList = make([]string, 0)
		for _, value := range classified[v] {
			userName, _ := value["name"].(string)
			classNameList = append(classNameList, userName)
		}
		rewriteList := strings.Join(classNameList, "\t")
		err = docx1.Replace(nameStr, rewriteList, -1)
		if err != nil {
			zap.L().Error(err.Error())
		}
	}
	err = docx1.Replace("{{Title}}", title, -1)
	if err != nil {
		zap.L().Error(err.Error())
	}
	filename := url.QueryEscape(title + ".docx")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename*=utf-8''%s", filename))
	err = docx1.Write(c.Writer)
	if err != nil {
		zap.L().Error(err.Error())
	}
	return
}

// func Demo(c *gin.Context) {
//
//	r, err := docx.ReadDocxFile("controller/routesJoinAudit/word/template.docx")
//	if err != nil {
//		log.Printf("docx Err %v", err)
//	}
//	defer func() {
//		_ = r.Close()
//	}()
//
//	// 使文档可编辑
//	docx1 := r.Editable()
//	// 替换模板中的占位符
//	err = docx1.Replace("{{LIST}}", "rewriteList\n asdfasdf", -1)
//	err = docx1.Replace("{{Title}}", "入团申请第十一期", -1)
//	if err != nil {
//		log.Printf("Replace Err %v", err)
//	}
//
//	c.Writer.Header().Set("Content-Disposition", `attachment; filename="tes.docx"`)
//
//	err = docx1.Write(c.Writer)
//	if err != nil {
//		log.Printf("Write Err %v", err)
//	}
//
// }
