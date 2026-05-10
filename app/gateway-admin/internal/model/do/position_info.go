// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionInfo is the golang structure of table position_info for DAO operations like Where/Data.
type PositionInfo struct {
	g.Meta    `orm:"table:position_info, do:true"`
	Id        any         //
	FileId    any         // 图片文件ID
	GoodsName any         //
	Link      any         //
	Sort      any         //
	GoodsId   any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
