// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FlashSaleResult is the golang structure of table flash_sale_result for DAO operations like Where/Data.
type FlashSaleResult struct {
	g.Meta     `orm:"table:flash_sale_result, do:true"`
	Id         any         // 秒杀结果ID
	ResultId   any         // 业务结果ID
	UserId     any         // 用户ID
	GoodsId    any         // 商品ID
	ActivityId any         // 活动ID
	OrderNo    any         // 秒杀订单号
	Status     any         // 状态 0处理中 1成功 2失败
	Message    any         // 处理消息
	PayAmount  any         // 支付金额，单位分
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
