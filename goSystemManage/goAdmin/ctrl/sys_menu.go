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

//对菜单接口进行编辑

//获取菜单管理接口
func GetMenuInfo(c *gin.Context) {
	var Id request.GetById
	id := checkStructById(Id.Id, c)
	err, menu := service.GetMenuInfoById(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取菜单信息失败：%v", err), c)
	} else {
		response.SuccessWithData(menu, c)
	}
}

//创建新的菜单
func CreateMenu(c *gin.Context) {
	var menu model.SysMenu
	checkMenuStruct(&menu, c)
	err := service.CreatMenu(menu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建菜单数据失败,%v", err), c)
	} else {
		response.SuccessWithMessage(fmt.Sprintf("创建菜单数据成功,%v", err), c)
	}
}

//更新菜单信息
func UpdateMenu(c *gin.Context) {
	var menu model.SysMenu
	checkMenuStruct(&menu, c)
	err := service.UpdateMenu(menu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新菜单数据失败,%v", err), c)
	} else {
		response.SuccessWithMessage(fmt.Sprintf("更新菜单数据成功,%v", err), c)
	}
}

//获取全部菜单
func GetMenuAll(c *gin.Context) {
	err, list := service.AllMenuList()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取全部菜单数据失败,%v", err), c)
	} else {
		response.SuccessWithData(list, c)
	}
}

//删除菜单
func DeleteMenu(c *gin.Context) {
	var Id request.GetById
	id := checkStructById(Id.Id, c)
	err := service.DeleteMenuById(id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除菜单数据失败,%v", err), c)
	} else {
		response.SuccessWithMessage(fmt.Sprintf("删除菜单数据成功,%v", err), c)
	}
}
func checkMenuStruct(menu *model.SysMenu, c *gin.Context) {
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除数据失败,%v", err), c)
	}
	menuVerify := util.Rules{
		"ParentId":  {util.NotEmpty()},
		"Name":      {util.NotEmpty()},
		"Path":      {util.NotEmpty()},
		"Desc":      {util.NotEmpty()},
		"Sort":      {util.NotEmpty()},
		"Component": {util.NotEmpty()},
	}
	menuVerifyErr := util.Verify(menu, menuVerify)
	if menuVerifyErr != nil {
		response.FailWithMessage(menuVerifyErr.Error(), c)
		return
	}
}
