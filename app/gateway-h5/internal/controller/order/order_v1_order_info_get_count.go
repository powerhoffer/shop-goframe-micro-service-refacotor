package order

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"
)

func (c *ControllerV1) OrderInfoGetCount(ctx context.Context, req *v1.OrderInfoGetCountReq) (res *v1.OrderInfoGetCountRes, err error) {
	userID, err := userIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	grpcRes, err := c.OrderInfoClient.GetCount(ctx, &order_info.OrderInfoGetCountReq{UserId: userID})
	if err != nil {
		return nil, err
	}
	return &v1.OrderInfoGetCountRes{
		Pending:   grpcRes.Pending,
		Shipping:  grpcRes.Shipping,
		Delivered: grpcRes.Delivered,
		Completed: grpcRes.Completed,
		AfterSale: grpcRes.AfterSale,
	}, nil
}
