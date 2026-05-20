// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FlashSaleResultDao is the data access object for the table flash_sale_result.
type FlashSaleResultDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  FlashSaleResultColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// FlashSaleResultColumns defines and stores column names for the table flash_sale_result.
type FlashSaleResultColumns struct {
	Id         string // 秒杀结果ID
	ResultId   string // 业务结果ID
	UserId     string // 用户ID
	GoodsId    string // 商品ID
	ActivityId string // 活动ID
	OrderNo    string // 秒杀订单号
	Status     string // 状态 0处理中 1成功 2失败
	Message    string // 处理消息
	PayAmount  string // 支付金额，单位分
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// flashSaleResultColumns holds the columns for the table flash_sale_result.
var flashSaleResultColumns = FlashSaleResultColumns{
	Id:         "id",
	ResultId:   "result_id",
	UserId:     "user_id",
	GoodsId:    "goods_id",
	ActivityId: "activity_id",
	OrderNo:    "order_no",
	Status:     "status",
	Message:    "message",
	PayAmount:  "pay_amount",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewFlashSaleResultDao creates and returns a new DAO object for table data access.
func NewFlashSaleResultDao(handlers ...gdb.ModelHandler) *FlashSaleResultDao {
	return &FlashSaleResultDao{
		group:    "default",
		table:    "flash_sale_result",
		columns:  flashSaleResultColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FlashSaleResultDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FlashSaleResultDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FlashSaleResultDao) Columns() FlashSaleResultColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FlashSaleResultDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FlashSaleResultDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FlashSaleResultDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
