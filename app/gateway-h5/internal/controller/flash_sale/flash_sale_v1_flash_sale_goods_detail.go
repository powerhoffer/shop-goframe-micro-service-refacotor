package flash_sale

import (
	"context"

	flashSaleGrpc "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/flash_sale/v1"
)

func (c *ControllerV1) FlashSaleGoodsDetail(ctx context.Context, req *v1.FlashSaleGoodsDetailReq) (res *v1.FlashSaleGoodsDetailRes, err error) {
	grpcRes, err := c.FlashSaleClient.GetFlashSaleGoodsDetail(ctx, &flashSaleGrpc.FlashSaleGoodsDetailReq{
		ActivityId: req.ActivityId,
		GoodsId:    req.GoodsId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.FlashSaleGoodsDetailRes{
		FlashSaleGoodsItem: buildGoodsItem(grpcRes.GoodsInfo),
	}, nil
}
