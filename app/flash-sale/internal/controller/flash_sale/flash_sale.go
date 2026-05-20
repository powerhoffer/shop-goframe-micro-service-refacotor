package flash_sale

import (
	"context"

	v1 "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	logic "shop-goframe-micro-service-refacotor/app/flash-sale/internal/logic/flash_sale"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedFlashSaleServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterFlashSaleServiceServer(s.Server, &Controller{})
}

func (*Controller) GetFlashSaleGoodsList(ctx context.Context, req *v1.FlashSaleGoodsListReq) (*v1.FlashSaleGoodsListRes, error) {
	return logic.GetGoodsList(ctx, req)
}

func (*Controller) GetFlashSaleGoodsDetail(ctx context.Context, req *v1.FlashSaleGoodsDetailReq) (*v1.FlashSaleGoodsDetailRes, error) {
	return logic.GetGoodsDetail(ctx, req)
}

func (*Controller) CreateFlashSaleOrder(ctx context.Context, req *v1.CreateFlashSaleOrderReq) (*v1.CreateFlashSaleOrderRes, error) {
	return logic.CreateOrder(ctx, req)
}

func (*Controller) GetFlashSaleResult(ctx context.Context, req *v1.GetFlashSaleResultReq) (*v1.GetFlashSaleResultRes, error) {
	return logic.GetResult(ctx, req)
}
