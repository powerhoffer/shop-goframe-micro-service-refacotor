// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FlashSaleOrderDao is the data access object for the table flash_sale_order.
type FlashSaleOrderDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  FlashSaleOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// FlashSaleOrderColumns defines and stores column names for the table flash_sale_order.
type FlashSaleOrderColumns struct {
	Id         string // 秒杀订单ID
	OrderNo    string // 秒杀订单号
	GoodsId    string // 商品ID
	ActivityId string // 活动ID
	UserId     string // 用户ID
	Count      string // 购买数量
	Amount     string // 实付金额，单位分
	Status     string // 状态 1成功 2失败 3取消
	ResultId   string // 秒杀结果ID
	Message    string // 处理消息
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// flashSaleOrderColumns holds the columns for the table flash_sale_order.
var flashSaleOrderColumns = FlashSaleOrderColumns{
	Id:         "id",
	OrderNo:    "order_no",
	GoodsId:    "goods_id",
	ActivityId: "activity_id",
	UserId:     "user_id",
	Count:      "count",
	Amount:     "amount",
	Status:     "status",
	ResultId:   "result_id",
	Message:    "message",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewFlashSaleOrderDao creates and returns a new DAO object for table data access.
func NewFlashSaleOrderDao(handlers ...gdb.ModelHandler) *FlashSaleOrderDao {
	return &FlashSaleOrderDao{
		group:    "default",
		table:    "flash_sale_order",
		columns:  flashSaleOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FlashSaleOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FlashSaleOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FlashSaleOrderDao) Columns() FlashSaleOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FlashSaleOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FlashSaleOrderDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FlashSaleOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
