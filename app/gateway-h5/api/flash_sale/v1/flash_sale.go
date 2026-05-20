package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 秒杀商品列表
type FlashSaleGoodsListReq struct {
	g.Meta     `path:"/flash-sale/goods" method:"get" tags:"秒杀管理" sm:"秒杀商品列表"`
	ActivityId uint32 `json:"activity_id" dc:"活动ID"`
	Page       uint32 `json:"page" d:"1" v:"min:1" dc:"页码"`
	Size       uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type FlashSaleGoodsListRes struct {
	List  []*FlashSaleGoodsItem `json:"list" dc:"秒杀商品列表"`
	Page  uint32                `json:"page" dc:"当前页码"`
	Size  uint32                `json:"size" dc:"每页数量"`
	Total uint32                `json:"total" dc:"总数"`
}

type FlashSaleGoodsItem struct {
	Id             uint64                 `json:"id" dc:"秒杀商品ID"`
	GoodsId        uint64                 `json:"goods_id" dc:"商品ID"`
	ActivityId     uint64                 `json:"activity_id" dc:"活动ID"`
	Title          string                 `json:"title" dc:"秒杀标题"`
	Description    string                 `json:"description" dc:"秒杀描述"`
	OriginalPrice  uint64                 `json:"original_price" dc:"原价，单位分"`
	SalePrice      uint64                 `json:"sale_price" dc:"秒杀价，单位分"`
	TotalStock     uint32                 `json:"total_stock" dc:"总库存"`
	AvailableStock uint32                 `json:"available_stock" dc:"可用库存"`
	StartTime      *timestamppb.Timestamp `json:"start_time" dc:"开始时间"`
	EndTime        *timestamppb.Timestamp `json:"end_time" dc:"结束时间"`
	Status         uint32                 `json:"status" dc:"状态 1启用 2禁用 3结束"`
	ImageUrl       string                 `json:"image_url" dc:"商品图片URL"`
	RemainSeconds  int64                  `json:"remain_seconds" dc:"剩余秒数"`
	CanBuy         bool                   `json:"can_buy" dc:"是否可购买"`
}

// 秒杀商品详情
type FlashSaleGoodsDetailReq struct {
	g.Meta     `path:"/flash-sale/goods/detail" method:"get" tags:"秒杀管理" sm:"秒杀商品详情"`
	ActivityId uint32 `json:"activity_id" v:"required" dc:"活动ID"`
	GoodsId    uint32 `json:"goods_id" v:"required" dc:"商品ID"`
}

type FlashSaleGoodsDetailRes struct {
	*FlashSaleGoodsItem
}

// 创建秒杀订单
type CreateFlashSaleOrderReq struct {
	g.Meta     `path:"/flash-sale/order" method:"post" tags:"秒杀管理" sm:"创建秒杀订单"`
	ActivityId uint32 `json:"activity_id" v:"required" dc:"活动ID"`
	GoodsId    uint32 `json:"goods_id" v:"required" dc:"商品ID"`
	Count      uint32 `json:"count" v:"required|min:1" dc:"购买数量"`
}

type CreateFlashSaleOrderRes struct {
	Success  bool   `json:"success" dc:"是否成功"`
	OrderNo  string `json:"order_no" dc:"秒杀订单号"`
	Message  string `json:"message" dc:"状态描述"`
	ResultId string `json:"result_id" dc:"结果查询ID"`
	Status   uint32 `json:"status" dc:"状态 0处理中 1成功 2失败"`
}

// 查询秒杀结果
type GetFlashSaleResultReq struct {
	g.Meta   `path:"/flash-sale/result" method:"get" tags:"秒杀管理" sm:"查询秒杀结果"`
	ResultId string `json:"result_id" v:"required" dc:"结果查询ID"`
}

type GetFlashSaleResultRes struct {
	Status    uint32 `json:"status" dc:"状态 0处理中 1成功 2失败"`
	Message   string `json:"message" dc:"状态描述"`
	OrderNo   string `json:"order_no" dc:"秒杀订单号"`
	GoodsId   uint32 `json:"goods_id" dc:"商品ID"`
	PayAmount uint64 `json:"pay_amount" dc:"支付金额，单位分"`
}
