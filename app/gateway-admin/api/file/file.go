// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package v1

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-admin/api/file/v1"
)

type IFileV1 interface {
	UploadImage(ctx context.Context, req *v1.UploadImageReq) (res *v1.UploadImageRes, err error)
}
