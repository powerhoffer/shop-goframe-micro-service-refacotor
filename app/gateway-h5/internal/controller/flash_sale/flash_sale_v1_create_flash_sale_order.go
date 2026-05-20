package flash_sale

import (
	"context"

	flashSaleGrpc "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/flash_sale/v1"
)

func (c *ControllerV1) CreateFlashSaleOrder(ctx context.Context, req *v1.CreateFlashSaleOrderReq) (res *v1.CreateFlashSaleOrderRes, err error) {
	userID, err := userIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	grpcRes, err := c.FlashSaleClient.CreateFlashSaleOrder(ctx, &flashSaleGrpc.CreateFlashSaleOrderReq{
		ActivityId: req.ActivityId,
		GoodsId:    req.GoodsId,
		UserId:     userID,
		Count:      req.Count,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateFlashSaleOrderRes{
		Success:  grpcRes.Success,
		OrderNo:  grpcRes.OrderNo,
		Message:  grpcRes.Message,
		ResultId: grpcRes.ResultId,
		Status:   grpcRes.Status,
	}, nil
}
