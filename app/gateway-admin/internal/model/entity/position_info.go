// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionInfo is the golang structure for table position_info.
type PositionInfo struct {
	Id        uint        `json:"id"        orm:"id"         description:""`       //
	FileId    int         `json:"fileId"    orm:"file_id"    description:"图片文件ID"` // 图片文件ID
	GoodsName string      `json:"goodsName" orm:"goods_name" description:""`       //
	Link      string      `json:"link"      orm:"link"       description:""`       //
	Sort      int         `json:"sort"      orm:"sort"       description:""`       //
	GoodsId   int         `json:"goodsId"   orm:"goods_id"   description:""`       //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`       //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`       //
}
