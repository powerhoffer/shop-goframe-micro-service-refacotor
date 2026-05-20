// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FlashSaleOrder is the golang structure of table flash_sale_order for DAO operations like Where/Data.
type FlashSaleOrder struct {
	g.Meta     `orm:"table:flash_sale_order, do:true"`
	Id         any         // 秒杀订单ID
	OrderNo    any         // 秒杀订单号
	GoodsId    any         // 商品ID
	ActivityId any         // 活动ID
	UserId     any         // 用户ID
	Count      any         // 购买数量
	Amount     any         // 实付金额，单位分
	Status     any         // 状态 1成功 2失败 3取消
	ResultId   any         // 秒杀结果ID
	Message    any         // 处理消息
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
