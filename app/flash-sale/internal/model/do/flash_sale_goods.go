// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FlashSaleGoods is the golang structure of table flash_sale_goods for DAO operations like Where/Data.
type FlashSaleGoods struct {
	g.Meta         `orm:"table:flash_sale_goods, do:true"`
	Id             any         // 秒杀商品ID
	GoodsId        any         // 商品ID
	ActivityId     any         // 活动ID
	Title          any         // 秒杀标题
	Description    any         // 秒杀描述
	OriginalPrice  any         // 原价，单位分
	SalePrice      any         // 秒杀价，单位分
	TotalStock     any         // 总库存
	AvailableStock any         // 可用库存
	StartTime      *gtime.Time // 开始时间
	EndTime        *gtime.Time // 结束时间
	Status         any         // 状态 1启用 2禁用 3结束
	ImageUrl       any         // 商品图片URL
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
