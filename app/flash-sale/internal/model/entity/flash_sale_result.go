// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FlashSaleResult is the golang structure for table flash_sale_result.
type FlashSaleResult struct {
	Id         uint64      `json:"id"         orm:"id"          description:"秒杀结果ID"`          // 秒杀结果ID
	ResultId   string      `json:"resultId"   orm:"result_id"   description:"业务结果ID"`          // 业务结果ID
	UserId     uint64      `json:"userId"     orm:"user_id"     description:"用户ID"`            // 用户ID
	GoodsId    uint64      `json:"goodsId"    orm:"goods_id"    description:"商品ID"`            // 商品ID
	ActivityId uint64      `json:"activityId" orm:"activity_id" description:"活动ID"`            // 活动ID
	OrderNo    string      `json:"orderNo"    orm:"order_no"    description:"秒杀订单号"`           // 秒杀订单号
	Status     uint        `json:"status"     orm:"status"      description:"状态 0处理中 1成功 2失败"` // 状态 0处理中 1成功 2失败
	Message    string      `json:"message"    orm:"message"     description:"处理消息"`            // 处理消息
	PayAmount  uint64      `json:"payAmount"  orm:"pay_amount"  description:"支付金额，单位分"`        // 支付金额，单位分
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`            // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`            // 更新时间
}
