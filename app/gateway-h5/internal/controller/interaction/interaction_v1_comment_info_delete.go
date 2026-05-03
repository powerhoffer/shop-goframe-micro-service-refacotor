package interaction

import (
	"context"
	comment "interaction/api/comment_info/v1"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/interaction/v1"
)

func (c *ControllerV1) CommentInfoDelete(ctx context.Context, req *v1.CommentInfoDeleteReq) (res *v1.CommentInfoDeleteRes, err error) {
	// 调用gRPC服务
	_, err = c.CommentInfoClient.Delete(ctx, &comment.CommentInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &v1.CommentInfoDeleteRes{}, nil
}
