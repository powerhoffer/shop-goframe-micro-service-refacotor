package flash_sale

import (
	"context"
	"fmt"
	"time"

	v1 "shop-goframe-micro-service-refacotor/app/flash-sale/api/flash_sale/v1"
	"shop-goframe-micro-service-refacotor/app/flash-sale/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/consts"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/dao"
	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/model/entity"
	"shop-goframe-micro-service-refacotor/app/flash-sale/utility/rabbitmq"
	flashRedis "shop-goframe-micro-service-refacotor/app/flash-sale/utility/redis"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type resultCache struct {
	Status    uint32 `json:"status"`
	Message   string `json:"message"`
	OrderNo   string `json:"order_no"`
	GoodsId   uint32 `json:"goods_id"`
	PayAmount uint64 `json:"pay_amount"`
}

func PreheatStock(ctx context.Context) error {
	var goodsList []*entity.FlashSaleGoods
	c := dao.FlashSaleGoods.Columns()
	if err := dao.FlashSaleGoods.Ctx(ctx).
		Where(c.Status, consts.GoodsStatusEnabled).
		WhereGT(c.EndTime, gtime.Now()).
		Scan(&goodsList); err != nil {
		return err
	}

	for _, item := range goodsList {
		if err := flashRedis.SetStock(ctx, uint32(item.ActivityId), uint32(item.GoodsId), item.AvailableStock); err != nil {
			return err
		}
	}
	g.Log().Infof(ctx, "秒杀库存预热完成，商品数量=%d", len(goodsList))
	return nil
}

func GetGoodsList(ctx context.Context, req *v1.FlashSaleGoodsListReq) (*v1.FlashSaleGoodsListRes, error) {
	page := int(req.PageNum)
	size := int(req.PageSize)
	if page <= 0 {
		page = 1
	}
	if size <= 0 || size > 50 {
		size = 10
	}

	c := dao.FlashSaleGoods.Columns()
	model := dao.FlashSaleGoods.Ctx(ctx)
	countModel := dao.FlashSaleGoods.Ctx(ctx)
	if req.ActivityId > 0 {
		model = model.Where(c.ActivityId, req.ActivityId)
		countModel = countModel.Where(c.ActivityId, req.ActivityId)
	}

	total, err := countModel.Count()
	if err != nil {
		return nil, err
	}

	var goodsList []*entity.FlashSaleGoods
	if err = model.OrderDesc(c.Id).Limit((page-1)*size, size).Scan(&goodsList); err != nil {
		return nil, err
	}

	res := &v1.FlashSaleGoodsListRes{
		Total: uint32(total),
		List:  make([]*v1.FlashSaleGoodsInfo, 0, len(goodsList)),
	}
	for _, item := range goodsList {
		res.List = append(res.List, buildGoodsInfo(ctx, item))
	}
	return res, nil
}

func GetGoodsDetail(ctx context.Context, req *v1.FlashSaleGoodsDetailReq) (*v1.FlashSaleGoodsDetailRes, error) {
	goods, err := getGoods(ctx, req.ActivityId, req.GoodsId)
	if err != nil {
		return nil, err
	}
	info := buildGoodsInfo(ctx, goods)
	return &v1.FlashSaleGoodsDetailRes{
		GoodsInfo:     info,
		RemainSeconds: info.RemainSeconds,
		CanBuy:        info.CanBuy,
	}, nil
}

func CreateOrder(ctx context.Context, req *v1.CreateFlashSaleOrderReq) (*v1.CreateFlashSaleOrderRes, error) {
	if req.UserId == 0 || req.GoodsId == 0 || req.ActivityId == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "user_id、goods_id、activity_id不能为空")
	}
	if req.Count == 0 {
		req.Count = 1
	}
	maxCount := g.Cfg().MustGet(ctx, "business.flashSale.maxCountPerOrder", 1).Uint32()
	if maxCount == 0 {
		maxCount = 1
	}
	if req.Count > maxCount {
		return failure("超过单次购买数量限制"), nil
	}

	if err := checkRateLimit(ctx, req.UserId); err != nil {
		return failure(err.Error()), nil
	}
	if purchased, err := isPurchased(ctx, req); err != nil {
		return nil, err
	} else if purchased {
		return failure("您已购买过该秒杀商品"), nil
	}

	goods, err := getGoods(ctx, req.ActivityId, req.GoodsId)
	if err != nil {
		return nil, err
	}
	if ok, msg := canBuy(goods); !ok {
		return failure(msg), nil
	}

	if err = flashRedis.ReduceStock(ctx, req.ActivityId, req.GoodsId, req.Count); err != nil {
		return failure(err.Error()), nil
	}

	resultId := newBusinessNo("FSR")
	orderNo := newBusinessNo("FSO")
	amount := uint64(req.Count) * goods.SalePrice
	if err = createOrderTransaction(ctx, req, goods, resultId, orderNo, amount); err != nil {
		flashRedis.IncreaseStock(ctx, req.ActivityId, req.GoodsId, req.Count)
		return nil, err
	}
	if err = flashRedis.SetPurchase(ctx, req.ActivityId, req.GoodsId, req.UserId, 24*time.Hour); err != nil {
		g.Log().Warningf(ctx, "写入秒杀限购缓存失败: %v", err)
	}

	cache := resultCache{
		Status:    consts.ResultStatusSuccess,
		Message:   "秒杀成功",
		OrderNo:   orderNo,
		GoodsId:   req.GoodsId,
		PayAmount: amount,
	}
	if err = flashRedis.SetJSON(ctx, flashRedis.ResultKey(resultId), cache, 24*time.Hour); err != nil {
		g.Log().Warningf(ctx, "写入秒杀结果缓存失败: %v", err)
	}

	if err = rabbitmq.PublishFlashSaleOrder(ctx, rabbitmq.FlashSaleOrderMessage{
		ResultId:   resultId,
		OrderNo:    orderNo,
		GoodsId:    req.GoodsId,
		ActivityId: req.ActivityId,
		UserId:     req.UserId,
		Count:      req.Count,
		Amount:     amount,
	}); err != nil {
		g.Log().Warningf(ctx, "发布秒杀订单消息失败: %v", err)
	}

	return &v1.CreateFlashSaleOrderRes{
		Success:  true,
		OrderNo:  orderNo,
		Message:  "秒杀成功",
		ResultId: resultId,
		Status:   consts.ResultStatusSuccess,
	}, nil
}

func GetResult(ctx context.Context, req *v1.GetFlashSaleResultReq) (*v1.GetFlashSaleResultRes, error) {
	if req.ResultId == "" || req.UserId == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "result_id、user_id不能为空")
	}

	var cache resultCache
	if ok, err := flashRedis.GetJSON(ctx, flashRedis.ResultKey(req.ResultId), &cache); err != nil {
		g.Log().Warningf(ctx, "读取秒杀结果缓存失败: %v", err)
	} else if ok {
		return &v1.GetFlashSaleResultRes{
			Status:    cache.Status,
			Message:   cache.Message,
			OrderNo:   cache.OrderNo,
			GoodsId:   cache.GoodsId,
			PayAmount: cache.PayAmount,
		}, nil
	}

	c := dao.FlashSaleResult.Columns()
	var result *entity.FlashSaleResult
	if err := dao.FlashSaleResult.Ctx(ctx).
		Where(c.ResultId, req.ResultId).
		Where(c.UserId, req.UserId).
		Scan(&result); err != nil {
		return nil, err
	}
	if result == nil {
		return &v1.GetFlashSaleResultRes{
			Status:  consts.ResultStatusProcessing,
			Message: "秒杀结果处理中",
		}, nil
	}
	return &v1.GetFlashSaleResultRes{
		Status:    uint32(result.Status),
		Message:   result.Message,
		OrderNo:   result.OrderNo,
		GoodsId:   uint32(result.GoodsId),
		PayAmount: result.PayAmount,
	}, nil
}

func getGoods(ctx context.Context, activityId, goodsId uint32) (*entity.FlashSaleGoods, error) {
	c := dao.FlashSaleGoods.Columns()
	var goods *entity.FlashSaleGoods
	err := dao.FlashSaleGoods.Ctx(ctx).
		Where(c.ActivityId, activityId).
		Where(c.GoodsId, goodsId).
		Scan(&goods)
	if err != nil {
		return nil, err
	}
	if goods == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "秒杀商品不存在")
	}
	return goods, nil
}

func buildGoodsInfo(ctx context.Context, goods *entity.FlashSaleGoods) *v1.FlashSaleGoodsInfo {
	pbGoods := &pbentity.FlashSaleGoods{}
	_ = gconv.Struct(goods, pbGoods)

	availableStock := uint32(goods.AvailableStock)
	if stock, err := flashRedis.GetStock(ctx, uint32(goods.ActivityId), uint32(goods.GoodsId)); err == nil && stock >= 0 {
		availableStock = uint32(stock)
		pbGoods.AvailableStock = availableStock
	}
	ok, _ := canBuy(goods)
	return &v1.FlashSaleGoodsInfo{
		Goods:          pbGoods,
		AvailableStock: availableStock,
		RemainSeconds:  remainSeconds(goods),
		CanBuy:         ok && availableStock > 0,
	}
}

func canBuy(goods *entity.FlashSaleGoods) (bool, string) {
	now := time.Now()
	if goods.Status != consts.GoodsStatusEnabled {
		return false, "秒杀商品未启用"
	}
	if goods.StartTime != nil && now.Before(goods.StartTime.Time) {
		return false, "秒杀尚未开始"
	}
	if goods.EndTime != nil && now.After(goods.EndTime.Time) {
		return false, "秒杀已结束"
	}
	if goods.AvailableStock <= 0 {
		return false, "库存不足"
	}
	return true, ""
}

func remainSeconds(goods *entity.FlashSaleGoods) int64 {
	if goods.EndTime == nil {
		return 0
	}
	remain := time.Until(goods.EndTime.Time)
	if remain <= 0 {
		return 0
	}
	return int64(remain.Seconds())
}

func checkRateLimit(ctx context.Context, userId uint32) error {
	if err := flashRedis.CheckLimit(ctx, "flash_sale:rate:global", 1000, time.Second); err != nil {
		return err
	}
	return flashRedis.CheckLimit(ctx, fmt.Sprintf("flash_sale:rate:user:%d", userId), 10, time.Second)
}

func isPurchased(ctx context.Context, req *v1.CreateFlashSaleOrderReq) (bool, error) {
	if ok, err := flashRedis.Exists(ctx, flashRedis.PurchaseKey(req.ActivityId, req.GoodsId, req.UserId)); err != nil {
		return false, err
	} else if ok {
		return true, nil
	}

	c := dao.FlashSaleOrder.Columns()
	count, err := dao.FlashSaleOrder.Ctx(ctx).
		Where(c.ActivityId, req.ActivityId).
		Where(c.GoodsId, req.GoodsId).
		Where(c.UserId, req.UserId).
		Count()
	return count > 0, err
}

func createOrderTransaction(ctx context.Context, req *v1.CreateFlashSaleOrderReq, goods *entity.FlashSaleGoods, resultId string, orderNo string, amount uint64) error {
	return dao.FlashSaleGoods.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := tx.Exec(
			"UPDATE flash_sale_goods SET available_stock = available_stock - ?, updated_at = NOW() WHERE id = ? AND available_stock >= ?",
			req.Count,
			goods.Id,
			req.Count,
		)
		if err != nil {
			return err
		}
		rows, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rows == 0 {
			return fmt.Errorf("库存不足")
		}

		now := gtime.Now()
		if _, err = tx.Insert(dao.FlashSaleOrder.Table(), g.Map{
			"order_no":    orderNo,
			"goods_id":    req.GoodsId,
			"activity_id": req.ActivityId,
			"user_id":     req.UserId,
			"count":       req.Count,
			"amount":      amount,
			"status":      consts.OrderStatusSuccess,
			"result_id":   resultId,
			"message":     "秒杀成功",
			"created_at":  now,
			"updated_at":  now,
		}); err != nil {
			return err
		}

		if _, err = tx.Insert(dao.FlashSaleResult.Table(), g.Map{
			"result_id":   resultId,
			"user_id":     req.UserId,
			"goods_id":    req.GoodsId,
			"activity_id": req.ActivityId,
			"order_no":    orderNo,
			"status":      consts.ResultStatusSuccess,
			"message":     "秒杀成功",
			"pay_amount":  amount,
			"created_at":  now,
			"updated_at":  now,
		}); err != nil {
			return err
		}

		return nil
	})
}

func failure(message string) *v1.CreateFlashSaleOrderRes {
	return &v1.CreateFlashSaleOrderRes{
		Success: false,
		Message: message,
		Status:  consts.ResultStatusFailed,
	}
}

func newBusinessNo(prefix string) string {
	return fmt.Sprintf("%s%s%s", prefix, time.Now().Format("20060102150405"), grand.Digits(8))
}
