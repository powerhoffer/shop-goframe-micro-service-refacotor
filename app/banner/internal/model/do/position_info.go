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
	PicUrl    any         // 图片链接
	GoodsName any         // 商品名称
	Link      any         // 跳转链接
	Sort      any         // 排序
	GoodsId   any         // 商品id
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
