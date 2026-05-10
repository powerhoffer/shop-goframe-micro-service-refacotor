// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileInfo is the golang structure of table file_info for DAO operations like Where/Data.
type FileInfo struct {
	g.Meta       `orm:"table:file_info, do:true"`
	Id           any         // 文件ID
	Name         any         // 文件名字
	Url          any         // 七牛云URL
	UploaderId   any         // 上传者ID（根据uploader_type区分是用户ID还是管理员ID）
	UploaderType any         // 上传者类型：1-H5用户，2-管理员
	FileType     any         // 文件类型：1-图片，2-视频，3-其他
	CreatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
}
