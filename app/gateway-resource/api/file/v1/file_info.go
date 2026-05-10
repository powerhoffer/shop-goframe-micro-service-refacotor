package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadImageReq struct {
	g.Meta `path:"/upload/image" tags:"文件上传" method:"post" summary:"上传图片" mime:"multipart/form-data"`
	// 注意：实际文件数据不会在JSON中传输 json - 是忽视json格式
	// 文件上传是通过 HTTP multipart/form-data 格式传输的 所以这里 json:"-"
	File         *ghttp.UploadFile `json:"file" type:"file" dc:"上传的文件" v:"required#请选择上传文件"`
	UploaderId   uint              `json:"uploader_id" type:"uint" dc:"上传者ID" v:"required#请选择上传者不能为空"`
	UploaderType uint              `json:"uploader_type" type:"uint" dc:"1-H5用户，2-管理员" v:"required#请选择上传者类型不能为空"`
	FileType     uint              `json:"file_type" type:"uint" dc:"1-图片，2-视频，3-其他" v:"required#请选择文件类型不能为空"`
}

type UploadImageRes struct {
	Url string `json:"url" dc:"图片访问URL"`
}
