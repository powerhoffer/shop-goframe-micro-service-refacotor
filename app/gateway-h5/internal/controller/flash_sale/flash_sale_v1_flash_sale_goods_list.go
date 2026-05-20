package flash_sale

import (
	"context"

	flashSaleGrpc "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"shop-goframe-micro-service-refacotor/app/gateway-h5/api/flash_sale/v1"
)

func (c *ControllerV1) FlashSaleGoodsList(ctx context.Context, req *v1.FlashSaleGoodsListReq) (res *v1.FlashSaleGoodsListRes, err error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 || req.Size > 100 {
		req.Size = 10
	}

	grpcRes, err := c.FlashSaleClient.GetFlashSaleGoodsList(ctx, &flashSaleGrpc.FlashSaleGoodsListReq{
		ActivityId: req.ActivityId,
		PageNum:    req.Page,
		PageSize:   req.Size,
	})
	if err != nil {
		return nil, err
	}

	res = &v1.FlashSaleGoodsListRes{
		Page:  req.Page,
		Size:  req.Size,
		Total: grpcRes.Total,
		List:  make([]*v1.FlashSaleGoodsItem, 0, len(grpcRes.List)),
	}
	for _, item := range grpcRes.List {
		if goods := buildGoodsItem(item); goods != nil {
			res.List = append(res.List, goods)
		}
	}
	return res, nil
}
