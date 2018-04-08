// UpdateGeographicInfo

package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type tmp_GeographicInfo struct {
	System_code string `form:"system_code" json:"system_code" binding:"required"`
	System_type int64
	Locat_long  string `form:"locat_long" json:"locat_long" binding:"required"` //经度
	Locat_lat   string `form:"locat_lat" json:"locat_lat" binding:"required"`   //纬度
}

func UpdateGeographicInfo(c *gin.Context) {
	tx, err := db.Beginx()
	if err != nil {
		cylog.Error(err)
		return
	}
	defer func() {

		err := recover()
		if err != nil {
			tx.Rollback()
			cylog.Error("有错误发生，正在回滚", err)

		} else {
			tx.Commit()
		}

	}()
	cylog.Info("==========================上传地理信息GPS开始:", c.Request.URL.Path, c.ClientIP())
	defer cylog.Info("==========================地理信息GPS结束\n\n")
	rtn := Out_Rtn{}
	var form tmp_GeographicInfo

	if err := c.ShouldBind(&form); err == nil {

		form.System_type = 7
		sqlcmd := ""
		sqlcmd = fmt.Sprintf(`insert into  tbase_locat(system_code,system_type,locat_long,locat_lat) values(
			:system_code,
			:system_type,
			:locat_long,
			:locat_lat
		)`)
		cylog.Debug(sqlcmd)
		if _, err = tx.NamedExec(sqlcmd, form); err != nil {
			rtn.Code = 401
			rtn.Message = "数据库内部错误,请重试"
			cylog.Error(err.Error())
			c.JSON(http.StatusOK, rtn)
			panic("cysql")
		}
		rtn.Code = 1
		rtn.Message = "OK"
		cylog.Info(rtn)
		c.JSON(http.StatusOK, rtn)
	} else {
		rtn.Code = 2
		rtn.Message = err.Error()
		cylog.Info(rtn)
		c.JSON(http.StatusOK, rtn)
	}

}
