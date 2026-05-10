// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RotationInfo is the golang structure for table rotation_info.
type RotationInfo struct {
	Id        uint        `json:"id"        orm:"id"         description:""`       //
	FileId    int         `json:"fileId"    orm:"file_id"    description:"图片文件ID"` // 图片文件ID
	Link      string      `json:"link"      orm:"link"       description:"跳转链接"`   // 跳转链接
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`     // 排序
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`       //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`       //
}
