// RiverNormal

package main

import (
	"fmt"

	"myfunc"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Out_Rtn struct {
	Code    int64
	Message string
}
type Tmp_tspss_dwd_all struct {
	Station_id   string `form:"station_id" binding:"required"`
	Guid_id      int64  `form:"guid"`
	Guid_id_10   int64
	Capture_time string `form:"capture_time" binding:"required"`
	Dwd_value    string `form:"dwd_value" binding:"required"`
}
type In_tspss_dwd_all struct {
	Station_id   string `form:"station_id" binding:"required"`
	Guid_id      string `form:"guid_id" binding:"required"`
	Capture_time string `form:"capture_time" binding:"required"`
	Dwd_value    string `form:"dwd_value" binding:"required"`
}

func RiverNormal(c *gin.Context) {
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

	cylog.Info("==========================上传普通信息开始:", c.Request.URL.Path, c.ClientIP())
	defer cylog.Info("==========================上传普通信息结束\n\n")

	cylog.Debugf("%+v\n", c.Request.Form)

	info_tmp := In_tspss_dwd_all{}
	rtn := Out_Rtn{}
	if err = c.BindJSON(&info_tmp); err != nil {
		cylog.Debug(err.Error())
		rtn.Code = 501
		rtn.Message = "json错误"
		c.JSON(http.StatusOK, rtn)
		return
	}
	cylog.Debugf("%+v\n", info_tmp)
	cylog.Debug(info_tmp.Capture_time)
	if len([]rune(info_tmp.Station_id)) > 8 || len([]rune(info_tmp.Station_id)) <= 0 {
		rtn.Code = 201
		rtn.Message = "station_id字段长度不可用"
		c.JSON(http.StatusOK, rtn)
		return
	}

	if len([]rune(info_tmp.Capture_time)) > 20 || len([]rune(info_tmp.Capture_time)) < 8 {
		rtn.Code = 203
		rtn.Message = "Capture_time字段长度不可用"
		c.JSON(http.StatusOK, rtn)
		return
	}
	if len([]rune(info_tmp.Dwd_value)) > 50 {
		rtn.Code = 204
		rtn.Message = "remark字段长度不可用"
		c.JSON(http.StatusOK, rtn)
		return
	}

	_, err = strconv.ParseInt(info_tmp.Station_id, 10, 64)
	if err != nil {
		rtn.Code = 302
		rtn.Message = "上传的Station_id_str为无效信息"
		c.JSON(http.StatusOK, rtn)
		return
	}

	_, err = time.Parse(myfunc.TimeLayout, info_tmp.Capture_time)
	if err != nil {
		rtn.Code = 304
		rtn.Message = "时间戳格式错误，请用：" + myfunc.TimeLayout
		c.JSON(http.StatusOK, rtn)
		return
	}
	oneItem := Tmp_tspss_dwd_all{}

	oneItem.Station_id = info_tmp.Station_id
	oneItem.Guid_id, err = strconv.ParseInt(info_tmp.Guid_id, 10, 64)
	if err != nil {
		rtn.Code = 301
		cylog.Info(err.Error())
		rtn.Message = "上传的guideID为无效信息"
		c.JSON(http.StatusOK, rtn)
		return
	}

	oneItem.Guid_id_10 = 0
	oneItem.Dwd_value = info_tmp.Dwd_value
	oneItem.Capture_time = info_tmp.Capture_time
	sqlcmd := ""
	if oneItem.Guid_id == 34 {
		sqlcmd = fmt.Sprintf(`insert into tspss_dwd_all(
		station_id    ,
		guid_id       ,
		guid_id_10    ,		
		capture_time  ,
		dwd_value         
		) values(
		:station_id    ,
		:guid_id       ,
		:guid_id_10    ,
		:capture_time  ,
		:dwd_value   
	)`)
	} else if oneItem.Guid_id == 57 {
		sqlcmd = fmt.Sprintf(`insert into tspss_dwd_all(
		station_id    ,
		guid_id       ,
		guid_id_10    ,		
		capture_time  ,
		dwd_value         
		) values(
		:station_id    ,
		:guid_id       ,
		:guid_id_10    ,
		:capture_time  ,
		:dwd_value   
	)`)
	} else {
		sqlcmd = fmt.Sprintf(`insert into tspss_dwd_all(
		station_id    ,
		guid_id       ,
		guid_id_10    ,		
		capture_time  ,
		dwd_value         
		) values(
		:station_id    ,
		:guid_id       ,
		:guid_id_10    ,
		:capture_time  ,
		:dwd_value     
	)`)
	}

	cylog.Debug(sqlcmd)
	if _, err = tx.NamedExec(sqlcmd, oneItem); err != nil {
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
	return
}
