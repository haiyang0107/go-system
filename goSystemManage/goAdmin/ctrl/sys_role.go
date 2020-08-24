package ctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goAdmin/global/response"
	"goAdmin/model"
	"goAdmin/service"
	"goAdmin/util"
)

//角色的Ctrl

//创建角色
func CreateRole(c *gin.Context) {
	var sysRole *model.SysRole
	checkRoleStruct(sysRole, c)
	err := service.CreateSysRole(sysRole)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("新增角色信息失败：%v", err), c)
	} else {
		response.SuccessWithMessage(fmt.Sprintf("新增角色信息成功：%v", err), c)
	}
}
func checkRoleStruct(sysRole *model.SysRole, c *gin.Context) {
	err := c.ShouldBindJSON(&sysRole)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("解析数据格式失败,%v", err), c)
	}
	sysRoleVerify := util.Rules{
		"RoleId":   {util.NotEmpty()},
		"RoleName": {util.NotEmpty()},
		"ParentId": {util.NotEmpty()},
	}
	sysRoleVerifyErr := util.Verify(sysRole, sysRoleVerify)
	if sysRoleVerifyErr != nil {
		response.FailWithMessage(sysRoleVerifyErr.Error(), c)
		return
	}
}
