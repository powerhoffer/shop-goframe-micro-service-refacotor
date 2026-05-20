package order

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) CancelOrder(ctx context.Context, req *v1.CancelOrderReq) (res *v1.CancelOrderRes, err error) {
	userID, err := userIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	grpcRes, err := c.OrderInfoClient.CancelOrder(ctx, &order_info.CancelOrderReq{
		Id:     req.Id,
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	res = &v1.CancelOrderRes{}
	if err := gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}
	return res, nil
}
