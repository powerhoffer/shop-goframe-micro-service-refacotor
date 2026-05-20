package flash_sale

import (
	"context"

	flashSaleGrpc "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/flash_sale/v1"
)

func (c *ControllerV1) GetFlashSaleResult(ctx context.Context, req *v1.GetFlashSaleResultReq) (res *v1.GetFlashSaleResultRes, err error) {
	userID, err := userIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	grpcRes, err := c.FlashSaleClient.GetFlashSaleResult(ctx, &flashSaleGrpc.GetFlashSaleResultReq{
		ResultId: req.ResultId,
		UserId:   userID,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetFlashSaleResultRes{
		Status:    grpcRes.Status,
		Message:   grpcRes.Message,
		OrderNo:   grpcRes.OrderNo,
		GoodsId:   grpcRes.GoodsId,
		PayAmount: grpcRes.PayAmount,
	}, nil
}
