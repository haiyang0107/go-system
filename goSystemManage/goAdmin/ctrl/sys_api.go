package ctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/base/request"
	response2 "goAdmin/base/response"
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
	id := checkStructById(ID.Id, c)
	err := service.DeleteApi(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除数据失败,%v", err), c)
	} else {
		response.SuccessWithMessage("删除成功", c)
	}
}
func checkStructById(ID int, c *gin.Context) (id int) {
	err := c.ShouldBindJSON(&ID)
	if err != nil {
		response.FailWithMessage("数据格式异常", c)
	}

	IdVerifyErr := util.Verify(ID, util.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
	}
	return ID
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

//获取api信息通过id
func GetById(c *gin.Context) {
	var ID request.GetById
	id := checkStructById(ID.Id, c)
	err, m := service.GetApiById(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败,%v", err), c)
	} else {
		response.SuccessWithData(m, c)
	}
}

//获取全部API信息,不分页
func AllListApi(c *gin.Context) {
	err, apis := service.GetAllListApi()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取全部API数据失败,%v", err), c)
	} else {
		response.SuccessWithData(apis, c)
	}
}

//分页查询-带搜索条件，
func PageListApi(c *gin.Context) {
	var params request.ApiParams
	c.ShouldBindJSON(params)
	PageVerifyErr := util.Verify(params.PageStrut, util.CustomizeMap["pageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
	}
	//service 获取 数据信息
	err, list, total := service.PageApisList(params.SysApi, params.PageStrut, params.Desc, params.OrderString)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取分页数据失败,%v", err), c)
	} else {
		response.FailWithData(response2.PageResult{
			List:     list,
			Total:    total,
			PageSize: params.PageSize,
			PageNum:  params.PageNum,
		}, c)
	}
}
