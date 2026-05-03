// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminInfo is the golang structure of table admin_info for DAO operations like Where/Data.
type AdminInfo struct {
	g.Meta    `orm:"table:admin_info, do:true"`
	Id        any         //
	Name      any         // 用户名
	Password  any         // 密码
	RoleIds   any         // 角色ids
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	UserSalt  any         // 加密盐
	IsAdmin   any         // 是否超级管理员
}
