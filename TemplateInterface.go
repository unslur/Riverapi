// TemplateInterface
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type tmpResult struct {
	Code    int64
	Message string
	Data    []tmpList
}
type tmpList struct {
	Product   string
	Prodfrede int64
}

func queryLastPlace(c *gin.Context) {

	defer func() {
		cylog.Flush()
		err := recover()
		if err != nil {
			cylog.Error("有错误发生，正在回滚", err)

		}

	}()
	cylog.Info("==========================查询GPS开始:", c.Request.URL.Path, c.ClientIP())
	defer cylog.Info("==========================查询GPS结束\n\n")
	CompanyCode := c.PostForm("Company_code")

	file, err := c.FormFile("file")
	if err != nil {
		cylog.Error(err.Error())
		return
	}
	rtn := tmpResult{}
	cylog.Info(file.Filename)
	if err = c.SaveUploadedFile(file, file.Filename); err != nil {
		cylog.Error(err.Error())
		rtn.Code = 2
		rtn.Message = err.Error()
		c.JSON(http.StatusOK, rtn)
		return
	}
	//cylog.Info(c.Request.Form)

	rtn.Code = 2
	rtn.Message = CompanyCode
	cylog.Info(rtn)
	c.JSON(http.StatusOK, rtn)
	return
}
