package order

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) OrderInfoGetDetail(ctx context.Context, req *v1.OrderInfoGetDetailReq) (res *v1.OrderInfoGetDetailRes, err error) {
	userID, err := userIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	grpcRes, err := c.OrderInfoClient.GetDetail(ctx, &order_info.OrderInfoGetDetailReq{
		Id:     req.Id,
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	res = &v1.OrderInfoGetDetailRes{}
	if grpcRes.OrderInfo != nil {
		res.OrderInfo = &v1.OrderInfoItem{}
		if err := gconv.Struct(grpcRes.OrderInfo, res.OrderInfo); err != nil {
			return nil, err
		}
	}
	if err := gconv.Structs(grpcRes.OrderGoodsInfos, &res.OrderGoodsInfos); err != nil {
		return nil, err
	}
	return res, nil
}
