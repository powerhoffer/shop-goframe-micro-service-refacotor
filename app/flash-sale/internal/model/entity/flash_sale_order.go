// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FlashSaleOrder is the golang structure for table flash_sale_order.
type FlashSaleOrder struct {
	Id         uint64      `json:"id"         orm:"id"          description:"秒杀订单ID"`         // 秒杀订单ID
	OrderNo    string      `json:"orderNo"    orm:"order_no"    description:"秒杀订单号"`          // 秒杀订单号
	GoodsId    uint64      `json:"goodsId"    orm:"goods_id"    description:"商品ID"`           // 商品ID
	ActivityId uint64      `json:"activityId" orm:"activity_id" description:"活动ID"`           // 活动ID
	UserId     uint64      `json:"userId"     orm:"user_id"     description:"用户ID"`           // 用户ID
	Count      uint        `json:"count"      orm:"count"       description:"购买数量"`           // 购买数量
	Amount     uint64      `json:"amount"     orm:"amount"      description:"实付金额，单位分"`       // 实付金额，单位分
	Status     uint        `json:"status"     orm:"status"      description:"状态 1成功 2失败 3取消"` // 状态 1成功 2失败 3取消
	ResultId   string      `json:"resultId"   orm:"result_id"   description:"秒杀结果ID"`         // 秒杀结果ID
	Message    string      `json:"message"    orm:"message"     description:"处理消息"`           // 处理消息
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`           // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`           // 更新时间
}
