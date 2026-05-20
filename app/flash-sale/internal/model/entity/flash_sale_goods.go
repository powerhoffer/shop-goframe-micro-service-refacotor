// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FlashSaleGoods is the golang structure for table flash_sale_goods.
type FlashSaleGoods struct {
	Id             uint64      `json:"id"             orm:"id"              description:"秒杀商品ID"`         // 秒杀商品ID
	GoodsId        uint64      `json:"goodsId"        orm:"goods_id"        description:"商品ID"`           // 商品ID
	ActivityId     uint64      `json:"activityId"     orm:"activity_id"     description:"活动ID"`           // 活动ID
	Title          string      `json:"title"          orm:"title"           description:"秒杀标题"`           // 秒杀标题
	Description    string      `json:"description"    orm:"description"     description:"秒杀描述"`           // 秒杀描述
	OriginalPrice  uint64      `json:"originalPrice"  orm:"original_price"  description:"原价，单位分"`         // 原价，单位分
	SalePrice      uint64      `json:"salePrice"      orm:"sale_price"      description:"秒杀价，单位分"`        // 秒杀价，单位分
	TotalStock     uint        `json:"totalStock"     orm:"total_stock"     description:"总库存"`            // 总库存
	AvailableStock uint        `json:"availableStock" orm:"available_stock" description:"可用库存"`           // 可用库存
	StartTime      *gtime.Time `json:"startTime"      orm:"start_time"      description:"开始时间"`           // 开始时间
	EndTime        *gtime.Time `json:"endTime"        orm:"end_time"        description:"结束时间"`           // 结束时间
	Status         uint        `json:"status"         orm:"status"          description:"状态 1启用 2禁用 3结束"` // 状态 1启用 2禁用 3结束
	ImageUrl       string      `json:"imageUrl"       orm:"image_url"       description:"商品图片URL"`        // 商品图片URL
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`           // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`           // 更新时间
}
