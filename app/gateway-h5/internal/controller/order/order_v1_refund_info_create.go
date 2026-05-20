package order

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	refund_info "shop-goframe-micro-service-refacotor/app/order/api/refund_info/v1"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) RefundInfoCreate(ctx context.Context, req *v1.RefundInfoCreateReq) (res *v1.RefundInfoCreateRes, err error) {
	grpcReq := &refund_info.RefundInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	userID, err := userIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	grpcReq.UserId = userID

	grpcRes, err := c.RefundInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.RefundInfoCreateRes{Id: grpcRes.Id}, nil
}
