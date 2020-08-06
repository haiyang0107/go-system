package ctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/base/request"
	resp "goAdmin/base/response"
	"goAdmin/global/response"
	"goAdmin/service"
	"goAdmin/util"
)

//对菜单接口进行编辑

//获取菜单分页信息
func GetMenuPage(c *gin.Context) {
	var params request.PageStrut
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("数据格式异常", c)
	}
	pageVerify := util.Verify(params, util.CustomizeMap["PageVerify"])
	if pageVerify != nil {
		response.FailWithMessage(pageVerify.Error(), c)
	}
	err, menuList, total := service.GetMenuPage(params)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取菜单信息失败：#{err}"), c)
	} else {
		response.SuccessWithData(resp.PageResult{
			List:     menuList,
			Total:    total,
			PageSize: params.PageSize,
			PageNum:  params.PageNum,
		}, c)
	}
}
