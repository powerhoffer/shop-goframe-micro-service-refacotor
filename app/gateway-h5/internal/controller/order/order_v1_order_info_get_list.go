package order

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) OrderInfoGetList(ctx context.Context, req *v1.OrderInfoGetListReq) (res *v1.OrderInfoGetListRes, err error) {
	grpcReq := &order_info.OrderInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	userID, err := userIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	grpcReq.UserId = userID

	grpcRes, err := c.OrderInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	if grpcRes == nil || grpcRes.Data == nil {
		return &v1.OrderInfoGetListRes{List: make([]*v1.OrderInfoItem, 0), Page: req.Page, Size: req.Size}, nil
	}

	res = &v1.OrderInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}
	if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}
	return res, nil
}
