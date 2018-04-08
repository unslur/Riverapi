// RiverNormalPost

package main

import (
	"fmt"

	"myfunc"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func RiverNormalPost(c *gin.Context) {
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
	Station_id_str := c.Query("station_id")     //站点ID
	Guid_id_str := c.Query("guid_id")           //项目ID
	Capture_time_str := c.Query("capture_time") //监测时间
	Dwd_value := c.Query("dwd_value")           //监测值
	//cylog.Debugf("%+v\n", c.Request.)
	//	cylog.Debug(Station_id_str)
	//	cylog.Debug(Guid_id_str)
	//	cylog.Debug(Capture_time_str)
	//	cylog.Debug(Dwd_value)
	info_tmp := Tmp_tspss_dwd_all{}
	rtn := Out_Rtn{}
	if c.BindJSON(&info_tmp) != nil {
		rtn.Code = 501
		rtn.Message = "json错误"
		c.JSON(http.StatusOK, rtn)
		return
	}
	cylog.Debugf("%+v\n", info_tmp)
	cylog.Debug(info_tmp.Capture_time)
	if len([]rune(Station_id_str)) > 8 || len([]rune(Station_id_str)) <= 0 {
		rtn.Code = 201
		rtn.Message = "station_id字段长度不可用"
		c.JSON(http.StatusOK, rtn)
		return
	}
	if len([]rune(Guid_id_str)) > 50 || len([]rune(Guid_id_str)) <= 0 {
		rtn.Code = 202
		rtn.Message = "guideID字段长度不可用"
		c.JSON(http.StatusOK, rtn)
		return
	}
	if len([]rune(Capture_time_str)) > 20 || len([]rune(Capture_time_str)) < 8 {
		rtn.Code = 203
		rtn.Message = "Capture_time字段长度不可用"
		c.JSON(http.StatusOK, rtn)
		return
	}
	if len([]rune(Dwd_value)) > 50 {
		rtn.Code = 204
		rtn.Message = "remark字段长度不可用"
		c.JSON(http.StatusOK, rtn)
		return
	}
	GuidId, err := strconv.ParseInt(Guid_id_str, 10, 64)
	if err != nil {
		rtn.Code = 301
		cylog.Info(err.Error())
		rtn.Message = "上传的guideID为无效信息"
		c.JSON(http.StatusOK, rtn)
		return
	}
	_, err = strconv.ParseInt(Station_id_str, 10, 64)
	if err != nil {
		rtn.Code = 302
		rtn.Message = "上传的Station_id_str为无效信息"
		c.JSON(http.StatusOK, rtn)
		return
	}
	//	Capture_time, err := strconv.ParseInt(Capture_time_str, 10, 64)
	//	if err != nil {
	//		rtn.Code = 303
	//		rtn.Message = "上传的guideID为无效信息"
	//		c.JSON(http.StatusOK, rtn)
	//		return
	//	}
	_, err = time.Parse(myfunc.TimeLayout, Capture_time_str)
	if err != nil {
		rtn.Code = 304
		rtn.Message = "时间戳格式错误，请用：" + myfunc.TimeLayout
		c.JSON(http.StatusOK, rtn)
		return
	}
	oneItem := Tmp_tspss_dwd_all{}

	oneItem.Station_id = Station_id_str
	oneItem.Guid_id = GuidId
	oneItem.Guid_id_10 = 0
	oneItem.Dwd_value = Dwd_value
	oneItem.Capture_time = Capture_time_str
	sqlcmd := fmt.Sprintf(`insert into tspss_dwd_all(
		station_id ,
		guid_id     ,
		guid_id_10  ,
		
		capture_time        ,
		dwd_value         
		) values(
		:station_id    ,
		:guid_id       ,
		:guid_id_10    ,
		:capture_time  ,
		:dwd_value     
	)`)
	cylog.Debug(sqlcmd)
	if _, err = tx.NamedExec(sqlcmd, oneItem); err != nil {
		rtn.Code = 401
		rtn.Message = "数据库内部错误，请重试"
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
