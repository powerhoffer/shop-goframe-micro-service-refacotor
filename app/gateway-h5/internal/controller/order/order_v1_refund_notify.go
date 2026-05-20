package order

import (
	"context"
	"io"
	"net/http"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	refund_info "shop-goframe-micro-service-refacotor/app/order/api/refund_info/v1"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) RefundNotify(ctx context.Context, req *v1.RefundNotifyReq) (res *v1.RefundNotifyRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		return nil, nil
	}

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		r.Response.WriteHeader(http.StatusBadRequest)
		r.Response.WriteJson(g.Map{"code": "FAIL", "message": "invalid request"})
		return nil, nil
	}

	_, err = c.RefundInfoClient.RefundNotify(ctx, &refund_info.RefundNotifyReq{
		RawBody: string(body),
		Headers: wechatHeaders(r),
	})
	if err != nil {
		g.Log().Errorf(ctx, "微信退款回调处理失败: %v", err)
		r.Response.WriteHeader(http.StatusInternalServerError)
		r.Response.WriteJson(g.Map{"code": "FAIL", "message": "network error"})
		return nil, nil
	}

	r.Response.WriteHeader(http.StatusOK)
	return nil, nil
}
