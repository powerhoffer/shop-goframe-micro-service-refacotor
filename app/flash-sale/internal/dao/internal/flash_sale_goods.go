// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FlashSaleGoodsDao is the data access object for the table flash_sale_goods.
type FlashSaleGoodsDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  FlashSaleGoodsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// FlashSaleGoodsColumns defines and stores column names for the table flash_sale_goods.
type FlashSaleGoodsColumns struct {
	Id             string // 秒杀商品ID
	GoodsId        string // 商品ID
	ActivityId     string // 活动ID
	Title          string // 秒杀标题
	Description    string // 秒杀描述
	OriginalPrice  string // 原价，单位分
	SalePrice      string // 秒杀价，单位分
	TotalStock     string // 总库存
	AvailableStock string // 可用库存
	StartTime      string // 开始时间
	EndTime        string // 结束时间
	Status         string // 状态 1启用 2禁用 3结束
	ImageUrl       string // 商品图片URL
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// flashSaleGoodsColumns holds the columns for the table flash_sale_goods.
var flashSaleGoodsColumns = FlashSaleGoodsColumns{
	Id:             "id",
	GoodsId:        "goods_id",
	ActivityId:     "activity_id",
	Title:          "title",
	Description:    "description",
	OriginalPrice:  "original_price",
	SalePrice:      "sale_price",
	TotalStock:     "total_stock",
	AvailableStock: "available_stock",
	StartTime:      "start_time",
	EndTime:        "end_time",
	Status:         "status",
	ImageUrl:       "image_url",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewFlashSaleGoodsDao creates and returns a new DAO object for table data access.
func NewFlashSaleGoodsDao(handlers ...gdb.ModelHandler) *FlashSaleGoodsDao {
	return &FlashSaleGoodsDao{
		group:    "default",
		table:    "flash_sale_goods",
		columns:  flashSaleGoodsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FlashSaleGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FlashSaleGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FlashSaleGoodsDao) Columns() FlashSaleGoodsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FlashSaleGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FlashSaleGoodsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FlashSaleGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
