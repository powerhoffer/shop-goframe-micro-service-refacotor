// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RotationInfo is the golang structure of table rotation_info for DAO operations like Where/Data.
type RotationInfo struct {
	g.Meta    `orm:"table:rotation_info, do:true"`
	Id        any         //
	PicUrl    any         // 轮播图片
	Link      any         // 跳转链接
	Sort      any         // 排序字段
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
