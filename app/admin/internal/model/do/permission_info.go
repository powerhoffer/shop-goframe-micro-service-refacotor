// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PermissionInfo is the golang structure of table permission_info for DAO operations like Where/Data.
type PermissionInfo struct {
	g.Meta    `orm:"table:permission_info, do:true"`
	Id        any         //
	Name      any         // 权限名称
	Path      any         // 路径
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
