package order

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) Payment(ctx context.Context, req *v1.PaymentReq) (res *v1.PaymentRes, err error) {
	if _, err := userIDFromCtx(ctx); err != nil {
		return nil, err
	}

	grpcReq := &order_info.PaymentReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.OrderInfoClient.Payment(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PaymentRes{}
	if err := gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}
	return res, nil
}
