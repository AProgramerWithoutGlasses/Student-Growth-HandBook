package JoinAudit

import "studentGrow/dao/mysql"

type RecList struct {
	True  []int `form:"true"`
	False []int `form:"false"`
}
type ResListWithIsPass struct {
	ID        int
	NowStatus string
}

func IsPassWithJSON(cr RecList, passType string) (resList []ResListWithIsPass) {
	resList = make([]ResListWithIsPass, 0)
	if len(cr.True) != 0 {
		for _, id := range cr.True {
			var resMsg ResListWithIsPass
			resMsg.ID = id
			updatedJoinAudit := mysql.IsPass(id, passType, "true")
			resMsg.NowStatus = updatedJoinAudit.ClassIsPass
			resList = append(resList, resMsg)
		}
	}
	if len(cr.False) != 0 {
		for _, id := range cr.False {
			var resMsg ResListWithIsPass
			resMsg.ID = id
			updatedJoinAudit := mysql.IsPass(id, passType, "false")
			resMsg.NowStatus = updatedJoinAudit.ClassIsPass
			resList = append(resList, resMsg)
		}
	}
	return
}
