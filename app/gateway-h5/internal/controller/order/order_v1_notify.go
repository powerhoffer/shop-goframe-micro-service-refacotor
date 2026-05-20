package order

import (
	"context"
	"io"
	"net/http"

	v1 "shop-goframe-micro-service-refacotor/app/gateway-h5/api/order/v1"
	order_info "shop-goframe-micro-service-refacotor/app/order/api/order_info/v1"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) Notify(ctx context.Context, req *v1.NotifyReq) (res *v1.NotifyRes, err error) {
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

	_, err = c.OrderInfoClient.Notify(ctx, &order_info.NotifyReq{
		RawBody: string(body),
		Headers: wechatHeaders(r),
	})
	if err != nil {
		g.Log().Errorf(ctx, "微信支付回调处理失败: %v", err)
		r.Response.WriteHeader(http.StatusInternalServerError)
		r.Response.WriteJson(g.Map{"code": "FAIL", "message": "network error"})
		return nil, nil
	}

	r.Response.WriteHeader(http.StatusOK)
	return nil, nil
}

func wechatHeaders(r *ghttp.Request) map[string]string {
	return map[string]string{
		"Wechatpay-Signature": r.Header.Get("Wechatpay-Signature"),
		"Wechatpay-Timestamp": r.Header.Get("Wechatpay-Timestamp"),
		"Wechatpay-Nonce":     r.Header.Get("Wechatpay-Nonce"),
		"Wechatpay-Serial":    r.Header.Get("Wechatpay-Serial"),
		"X-Bypass-Verify":     r.Header.Get("X-Bypass-Verify"),
	}
}
