package ctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/base/request"
	"goAdmin/global/response"
	"goAdmin/model"
	"goAdmin/service"
	"goAdmin/util"
)

//创建api
func CreateApi(c *gin.Context) {
	var api model.SysApi
	checkStruct(api, c)
	error := service.CreateApi(api)
	if error != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败,%v", error), c)
	} else {
		response.SuccessWithMessage("创建成功", c)
	}
}

//更新api
func UpdateApi(c *gin.Context) {
	var api model.SysApi
	checkStruct(api, c)
	err := service.UpdateApi(api)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改数据失败,%v", err), c)
	} else {
		response.SuccessWithMessage("修改数据成功", c)
	}
}

//删除api
func DeleteApi(c *gin.Context) {
	var ID request.GetById
	err := c.ShouldBindJSON(&ID)
	if err != nil {
		response.FailWithMessage("数据格式异常", c)
	}

	IdVerifyErr := util.Verify(ID, util.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
	}
	err = service.DeleteApi(ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除数据失败,%v", err), c)
	} else {
		response.SuccessWithMessage("删除成功", c)
	}
}
func checkStruct(api model.SysApi, c *gin.Context) {
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage("数据格式异常", c)
	}
	apiVerify := util.Rules{
		"Path":        {util.NotEmpty()},
		"Group":       {util.NotEmpty()},
		"Method":      {util.NotEmpty()},
		"Description": {util.NotEmpty()},
	}
	apiVerifyErr := util.Verify(api, apiVerify)
	if apiVerifyErr != nil {
		response.FailWithMessage(apiVerifyErr.Error(), c)
		return
	}
}
