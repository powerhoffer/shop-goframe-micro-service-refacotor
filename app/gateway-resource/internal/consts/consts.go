package consts

const (
	UploadImageFail = "UploadImage 图片上传失败"
	FileInfo        = "FileInfo"

	GetListFail   = "GetList 查询失败"
	GetDetailFail = "GetDetail 查询失败"
	CreateFail    = "Create 插入失败"
	UpdateFail    = "Update 更新失败"
	DeleteFail    = "Delete 删除失败"
	LoginFail     = "Login 登录失败"
	RegisterFail  = "Register 注册失败"
)

func InfoError(info string, fail string) string {
	return info + " " + fail
}
