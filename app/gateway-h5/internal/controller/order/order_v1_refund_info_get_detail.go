package order

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	refund_info "shop-goframe-micro-service-refacotor/app/order/api/refund_info/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) RefundInfoGetDetail(ctx context.Context, req *v1.RefundInfoGetDetailReq) (res *v1.RefundInfoGetDetailRes, err error) {
	grpcRes, err := c.RefundInfoClient.GetDetail(ctx, &refund_info.RefundInfoGetDetailReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	if grpcRes == nil || grpcRes.Data == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "退款记录不存在")
	}

	res = &v1.RefundInfoGetDetailRes{RefundInfoItem: &v1.RefundInfoItem{}}
	if err := gconv.Struct(grpcRes.Data, res.RefundInfoItem); err != nil {
		return nil, err
	}
	return res, nil
}
