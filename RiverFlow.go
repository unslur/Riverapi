// RiverFlow

package main

import (
	"fmt"

	"myfunc"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

type Tmp_tspss_dwd_flow struct {
	Station_id       string `db:"station_id"`
	Capture_time     string `db:"capture_time"`
	Area             string `db:"area"`
	Waterlevel_start string `db:"waterlevel_start"`
	Waterlevel_end   string `db:"waterlevel_end"`
	Velocity         string `db:"velocity"`
	Leftside         string `db:"leftside"`
	Rightside        string `db:"rightside"`
	Flows            string `db:"flows"`
	Line_count       string `db:"line_count"`
	Linestartpos_1   string `db:"linestartpos_1"`
	Linedepth_1      string `db:"linedepth_1"`
	Linebottom_1     string `db:"linebottom_1"`
	Linevelocity_1   string `db:"linevelocity_1"`
	Subavrliusu_1    string `db:"subavrliusu_1"`
	Subarea_1        string `db:"subarea_1"`
	Subflows_1       string `db:"subflows_1"`
	Linestartpos_2   string `db:"linestartpos_2"`
	Linedepth_2      string `db:"linedepth_2"`
	Linebottom_2     string `db:"linebottom_2"`
	Linevelocity_2   string `db:"linevelocity_2"`
	Subavrliusu_2    string `db:"subavrliusu_2"`
	Subarea_2        string `db:"subarea_2"`
	Subflows_2       string `db:"subflows_2"`
	Linestartpos_3   string `db:"linestartpos_3"`
	Linedepth_3      string `db:"linedepth_3"`
	Linebottom_3     string `db:"linebottom_3"`
	Linevelocity_3   string `db:"linevelocity_3"`
	Subavrliusu_3    string `db:"subavrliusu_3"`
	Subarea_3        string `db:"subarea_3"`
	Subflows_3       string `db:"subflows_3"`
	Linestartpos_4   string `db:"linestartpos_4"`
	Linedepth_4      string `db:"linedepth_4"`
	Linebottom_4     string `db:"linebottom_4"`
	Linevelocity_4   string `db:"linevelocity_4"`
	Subavrliusu_4    string `db:"subavrliusu_4"`
	Subarea_4        string `db:"subarea_4"`
	Subflows_4       string `db:"subflows_4"`
	Linestartpos_5   string `db:"linestartpos_5"`
	Linedepth_5      string `db:"linedepth_5"`
	Linebottom_5     string `db:"linebottom_5"`
	Linevelocity_5   string `db:"linevelocity_5"`
	Subavrliusu_5    string `db:"subavrliusu_5"`
	Subarea_5        string `db:"subarea_5"`
	Subflows_5       string `db:"subflows_5"`
	Linestartpos_6   string `db:"linestartpos_6"`
	Linedepth_6      string `db:"linedepth_6"`
	Linebottom_6     string `db:"linebottom_6"`
	Linevelocity_6   string `db:"linevelocity_6"`
	Subavrliusu_6    string `db:"subavrliusu_6"`
	Subarea_6        string `db:"subarea_6"`
	Subflows_6       string `db:"subflows_6"`
	Linestartpos_7   string `db:"linestartpos_7"`
	Linedepth_7      string `db:"linedepth_7"`
	Linebottom_7     string `db:"linebottom_7"`
	Linevelocity_7   string `db:"linevelocity_7"`
	Subavrliusu_7    string `db:"subavrliusu_7"`
	Subarea_7        string `db:"subarea_7"`
	Subflows_7       string `db:"subflows_7"`
	Linestartpos_8   string `db:"linestartpos_8"`
	Linedepth_8      string `db:"linedepth_8"`
	Linebottom_8     string `db:"linebottom_8"`
	Linevelocity_8   string `db:"linevelocity_8"`
	Subavrliusu_8    string `db:"subavrliusu_8"`
	Subarea_8        string `db:"subarea_8"`
	Subflows_8       string `db:"subflows_8"`
	Linestartpos_9   string `db:"linestartpos_9"`
	Linedepth_9      string `db:"linedepth_9"`
	Linebottom_9     string `db:"linebottom_9"`
	Linevelocity_9   string `db:"linevelocity_9"`
	Subavrliusu_9    string `db:"subavrliusu_9"`
	Subarea_9        string `db:"subarea_9"`
	Subflows_9       string `db:"subflows_9"`
	Linestartpos_10  string `db:"linestartpos_10"`
	Linedepth_10     string `db:"linedepth_10"`
	Linebottom_10    string `db:"linebottom_10"`
	Linevelocity_10  string `db:"linevelocity_10"`
	Subavrliusu_10   string `db:"subavrliusu_10"`
	Subarea_10       string `db:"subarea_10"`
	Subflows_10      string `db:"subflows_10"`
	Linestartpos_11  string `db:"linestartpos_11"`
	Linedepth_11     string `db:"linedepth_11"`
	Linebottom_11    string `db:"linebottom_11"`
	Linevelocity_11  string `db:"linevelocity_11"`
	Subavrliusu_11   string `db:"subavrliusu_11"`
	Subarea_11       string `db:"subarea_11"`
	Subflows_11      string `db:"subflows_11"`
	Linestartpos_12  string `db:"linestartpos_12"`
	Linedepth_12     string `db:"linedepth_12"`
	Linebottom_12    string `db:"linebottom_12"`
	Linevelocity_12  string `db:"linevelocity_12"`
	Subavrliusu_12   string `db:"subavrliusu_12"`
	Subarea_12       string `db:"subarea_12"`
	Subflows_12      string `db:"subflows_12"`
	Linestartpos_13  string `db:"linestartpos_13"`
	Linedepth_13     string `db:"linedepth_13"`
	Linebottom_13    string `db:"linebottom_13"`
	Linevelocity_13  string `db:"linevelocity_13"`
	Subavrliusu_13   string `db:"subavrliusu_13"`
	Subarea_13       string `db:"subarea_13"`
	Subflows_13      string `db:"subflows_13"`
	Linestartpos_14  string `db:"linestartpos_14"`
	Linedepth_14     string `db:"linedepth_14"`
	Linebottom_14    string `db:"linebottom_14"`
	Linevelocity_14  string `db:"linevelocity_14"`
	Subavrliusu_14   string `db:"subavrliusu_14"`
	Subarea_14       string `db:"subarea_14"`
	Subflows_14      string `db:"subflows_14"`
	Linestartpos_15  string `db:"linestartpos_15"`
	Linedepth_15     string `db:"linedepth_15"`
	Linebottom_15    string `db:"linebottom_15"`
	Linevelocity_15  string `db:"linevelocity_15"`
	Subavrliusu_15   string `db:"subavrliusu_15"`
	Subarea_15       string `db:"subarea_15"`
	Subflows_15      string `db:"subflows_15"`
	Linestartpos_16  string `db:"linestartpos_16"`
	Linedepth_16     string `db:"linedepth_16"`
	Linebottom_16    string `db:"linebottom_16"`
	Linevelocity_16  string `db:"linevelocity_16"`
	Subavrliusu_16   string `db:"subavrliusu_16"`
	Subarea_16       string `db:"subarea_16"`
	Subflows_16      string `db:"subflows_16"`
	Linestartpos_17  string `db:"linestartpos_17"`
	Linedepth_17     string `db:"linedepth_17"`
	Linebottom_17    string `db:"linebottom_17"`
	Linevelocity_17  string `db:"linevelocity_17"`
	Subavrliusu_17   string `db:"subavrliusu_17"`
	Subarea_17       string `db:"subarea_17"`
	Subflows_17      string `db:"subflows_17"`
	Linestartpos_18  string `db:"linestartpos_18"`
	Linedepth_18     string `db:"linedepth_18"`
	Linebottom_18    string `db:"linebottom_18"`
	Linevelocity_18  string `db:"linevelocity_18"`
	Subavrliusu_18   string `db:"subavrliusu_18"`
	Subarea_18       string `db:"subarea_18"`
	Subflows_18      string `db:"subflows_18"`
	Linestartpos_19  string `db:"linestartpos_19"`
	Linedepth_19     string `db:"linedepth_19"`
	Linebottom_19    string `db:"linebottom_19"`
	Linevelocity_19  string `db:"linevelocity_19"`
	Subavrliusu_19   string `db:"subavrliusu_19"`
	Subarea_19       string `db:"subarea_19"`
	Subflows_19      string `db:"subflows_19"`
	Linestartpos_20  string `db:"linestartpos_20"`
	Linedepth_20     string `db:"linedepth_20"`
	Linebottom_20    string `db:"linebottom_20"`
	Linevelocity_20  string `db:"linevelocity_20"`
	Subavrliusu_20   string `db:"subavrliusu_20"`
	Subarea_20       string `db:"subarea_20"`
	Subflows_20      string `db:"subflows_20"`
	Linestartpos_21  string `db:"linestartpos_21"`
	Linedepth_21     string `db:"linedepth_21"`
	Linebottom_21    string `db:"linebottom_21"`
	Linevelocity_21  string `db:"linevelocity_21"`
	Subavrliusu_21   string `db:"subavrliusu_21"`
	Subarea_21       string `db:"subarea_21"`
	Subflows_21      string `db:"subflows_21"`
	Linestartpos_22  string `db:"linestartpos_22"`
	Linedepth_22     string `db:"linedepth_22"`
	Linebottom_22    string `db:"linebottom_22"`
	Linevelocity_22  string `db:"linevelocity_22"`
	Subavrliusu_22   string `db:"subavrliusu_22"`
	Subarea_22       string `db:"subarea_22"`
	Subflows_22      string `db:"subflows_22"`
	Linestartpos_23  string `db:"linestartpos_23"`
	Linedepth_23     string `db:"linedepth_23"`
	Linebottom_23    string `db:"linebottom_23"`
	Linevelocity_23  string `db:"linevelocity_23"`
	Subavrliusu_23   string `db:"subavrliusu_23"`
	Subarea_23       string `db:"subarea_23"`
	Subflows_23      string `db:"subflows_23"`
	Linestartpos_24  string `db:"linestartpos_24"`
	Linedepth_24     string `db:"linedepth_24"`
	Linebottom_24    string `db:"linebottom_24"`
	Linevelocity_24  string `db:"linevelocity_24"`
	Subavrliusu_24   string `db:"subavrliusu_24"`
	Subarea_24       string `db:"subarea_24"`
	Subflows_24      string `db:"subflows_24"`
}

func RiverFlow(c *gin.Context) {
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

	cylog.Info("==========================上传水流信息开始:", c.Request.URL.Path, c.ClientIP())
	defer cylog.Info("==========================上传水流信息结束\n\n")

	cylog.Debugf("%+v\n", c.Request.Form)

	info_tmp := Tmp_tspss_dwd_flow{}
	rtn := Out_Rtn{}
	if err = c.BindJSON(&info_tmp); err != nil {
		cylog.Debug(err.Error())
		rtn.Code = 501
		rtn.Message = "json错误"
		c.JSON(http.StatusOK, rtn)
		return
	}
	cylog.Debugf("%+v\n", info_tmp)
	_, err = time.Parse(myfunc.TimeLayout, info_tmp.Capture_time)
	if err != nil {
		rtn.Code = 304
		rtn.Message = "时间戳格式错误，请用：" + myfunc.TimeLayout
		c.JSON(http.StatusOK, rtn)
		return
	}
	sqlcmd := ""

	sqlcmd = fmt.Sprintf(`insert into tspss_dwd_flow(
		station_id,
capture_time,
area,
waterlevel_start,
waterlevel_end,
velocity,
leftside,
rightside,
flows,
line_count,
linestartpos_1,
linedepth_1,
linebottom_1,
linevelocity_1,
subavrliusu_1,
subarea_1,
subflows_1,
linestartpos_2,
linedepth_2,
linebottom_2,
linevelocity_2,
subavrliusu_2,
subarea_2,
subflows_2,
linestartpos_3,
linedepth_3,
linebottom_3,
linevelocity_3,
subavrliusu_3,
subarea_3,
subflows_3,
linestartpos_4,
linedepth_4,
linebottom_4,
linevelocity_4,
subavrliusu_4,
subarea_4,
subflows_4,
linestartpos_5,
linedepth_5,
linebottom_5,
linevelocity_5,
subavrliusu_5,
subarea_5,
subflows_5,
linestartpos_6,
linedepth_6,
linebottom_6,
linevelocity_6,
subavrliusu_6,
subarea_6,
subflows_6,
linestartpos_7,
linedepth_7,
linebottom_7,
linevelocity_7,
subavrliusu_7,
subarea_7,
subflows_7,
linestartpos_8,
linedepth_8,
linebottom_8,
linevelocity_8,
subavrliusu_8,
subarea_8,
subflows_8,
linestartpos_9,
linedepth_9,
linebottom_9,
linevelocity_9,
subavrliusu_9,
subarea_9,
subflows_9,
linestartpos_10,
linedepth_10,
linebottom_10,
linevelocity_10,
subavrliusu_10,
subarea_10,
subflows_10,
linestartpos_11,
linedepth_11,
linebottom_11,
linevelocity_11,
subavrliusu_11,
subarea_11,
subflows_11,
linestartpos_12,
linedepth_12,
linebottom_12,
linevelocity_12,
subavrliusu_12,
subarea_12,
subflows_12,
linestartpos_13,
linedepth_13,
linebottom_13,
linevelocity_13,
subavrliusu_13,
subarea_13,
subflows_13,
linestartpos_14,
linedepth_14,
linebottom_14,
linevelocity_14,
subavrliusu_14,
subarea_14,
subflows_14,
linestartpos_15,
linedepth_15,
linebottom_15,
linevelocity_15,
subavrliusu_15,
subarea_15,
subflows_15,
linestartpos_16,
linedepth_16,
linebottom_16,
linevelocity_16,
subavrliusu_16,
subarea_16,
subflows_16,
linestartpos_17,
linedepth_17,
linebottom_17,
linevelocity_17,
subavrliusu_17,
subarea_17,
subflows_17,
linestartpos_18,
linedepth_18,
linebottom_18,
linevelocity_18,
subavrliusu_18,
subarea_18,
subflows_18,
linestartpos_19,
linedepth_19,
linebottom_19,
linevelocity_19,
subavrliusu_19,
subarea_19,
subflows_19,
linestartpos_20,
linedepth_20,
linebottom_20,
linevelocity_20,
subavrliusu_20,
subarea_20,
subflows_20,
linestartpos_21,
linedepth_21,
linebottom_21,
linevelocity_21,
subavrliusu_21,
subarea_21,
subflows_21,
linestartpos_22,
linedepth_22,
linebottom_22,
linevelocity_22,
subavrliusu_22,
subarea_22,
subflows_22,
linestartpos_23,
linedepth_23,
linebottom_23,
linevelocity_23,
subavrliusu_23,
subarea_23,
subflows_23,
linestartpos_24,
linedepth_24,
linebottom_24,
linevelocity_24,
subavrliusu_24,
subarea_24,
subflows_24
    
		) values(
		:station_id,
	:capture_time,
	:areas,
	:waterlevel_start,
	:waterlevel_end,
	:velocity,
	:leftside,
	:rightside,
	:flows,
	:line_count,
	:linestartpos_1,
	:linedepth_1,
	:linebottom_1,
	:linevelocity_1,
	:subavrliusu_1,
	:subarea_1,
	:subflows_1,
	:linestartpos_2,
	:linedepth_2,
	:linebottom_2,
	:linevelocity_2,
	:subavrliusu_2,
	:subarea_2,
	:subflows_2,
	:linestartpos_3,
	:linedepth_3,
	:linebottom_3,
	:linevelocity_3,
	:subavrliusu_3,
	:subarea_3,
	:subflows_3,
	:linestartpos_4,
	:linedepth_4,
	:linebottom_4,
	:linevelocity_4,
	:subavrliusu_4,
	:subarea_4,
	:subflows_4,
	:linestartpos_5,
	:linedepth_5,
	:linebottom_5,
	:linevelocity_5,
	:subavrliusu_5,
	:subarea_5,
	:subflows_5,
	:linestartpos_6,
	:linedepth_6,
	:linebottom_6,
	:linevelocity_6,
	:subavrliusu_6,
	:subarea_6,
	:subflows_6,
	:linestartpos_7,
	:linedepth_7,
	:linebottom_7,
	:linevelocity_7,
	:subavrliusu_7,
	:subarea_7,
	:subflows_7,
	:linestartpos_8,
	:linedepth_8,
	:linebottom_8,
	:linevelocity_8,
	:subavrliusu_8,
	:subarea_8,
	:subflows_8,
	:linestartpos_9,
	:linedepth_9,
	:linebottom_9,
	:linevelocity_9,
	:subavrliusu_9,
	:subarea_9,
	:subflows_9,
	:linestartpos_10,
	:linedepth_10,
	:linebottom_10,
	:linevelocity_10,
	:subavrliusu_10,
	:subarea_10,
	:subflows_10,
	:linestartpos_11,
	:linedepth_11,
	:linebottom_11,
	:linevelocity_11,
	:subavrliusu_11,
	:subarea_11,
	:subflows_11,
	:linestartpos_12,
	:linedepth_12,
	:linebottom_12,
	:linevelocity_12,
	:subavrliusu_12,
	:subarea_12,
	:subflows_12,
	:linestartpos_13,
	:linedepth_13,
	:linebottom_13,
	:linevelocity_13,
	:subavrliusu_13,
	:subarea_13,
	:subflows_13,
	:linestartpos_14,
	:linedepth_14,
	:linebottom_14,
	:linevelocity_14,
	:subavrliusu_14,
	:subarea_14,
	:subflows_14,
	:linestartpos_15,
	:linedepth_15,
	:linebottom_15,
	:linevelocity_15,
	:subavrliusu_15,
	:subarea_15,
	:subflows_15,
	:linestartpos_16,
	:linedepth_16,
	:linebottom_16,
	:linevelocity_16,
	:subavrliusu_16,
	:subarea_16,
	:subflows_16,
	:linestartpos_17,
	:linedepth_17,
	:linebottom_17,
	:linevelocity_17,
	:subavrliusu_17,
	:subarea_17,
	:subflows_17,
	:linestartpos_18,
	:linedepth_18,
	:linebottom_18,
	:linevelocity_18,
	:subavrliusu_18,
	:subarea_18,
	:subflows_18,
	:linestartpos_19,
	:linedepth_19,
	:linebottom_19,
	:linevelocity_19,
	:subavrliusu_19,
	:subarea_19,
	:subflows_19,
	:linestartpos_20,
	:linedepth_20,
	:linebottom_20,
	:linevelocity_20,
	:subavrliusu_20,
	:subarea_20,
	:subflows_20,
	:linestartpos_21,
	:linedepth_21,
	:linebottom_21,
	:linevelocity_21,
	:subavrliusu_21,
	:subarea_21,
	:subflows_21,
	:linestartpos_22,
	:linedepth_22,
	:linebottom_22,
	:linevelocity_22,
	:subavrliusu_22,
	:subarea_22,
	:subflows_22,
	:linestartpos_23,
	:linedepth_23,
	:linebottom_23,
	:linevelocity_23,
	:subavrliusu_23,
	:subarea_23,
	:subflows_23,
	:linestartpos_24,
	:linedepth_24,
	:linebottom_24,
	:linevelocity_24,
	:subavrliusu_24,
	:subarea_24,
	:subflows_24       
	)`)

	cylog.Debug(sqlcmd)
	if _, err = tx.NamedExec(sqlcmd, info_tmp); err != nil {
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
