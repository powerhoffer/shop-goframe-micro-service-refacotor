package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

var client *gredis.Redis

func Init(ctx context.Context) error {
	for _, key := range []string{"redis.flash_sale", "redis.goods", "redis.default"} {
		cfgVar, err := g.Cfg().Get(ctx, key)
		if err != nil || cfgVar.IsNil() {
			continue
		}

		cfg := &gredis.Config{}
		if err = cfgVar.Scan(cfg); err != nil {
			return fmt.Errorf("解析Redis配置失败: %w", err)
		}
		redis, err := gredis.New(cfg)
		if err != nil {
			return fmt.Errorf("创建Redis连接失败: %w", err)
		}
		if _, err = redis.Do(ctx, "PING"); err != nil {
			return fmt.Errorf("Redis连接测试失败: %w", err)
		}
		client = redis
		g.Log().Infof(ctx, "秒杀Redis初始化成功: %s", key)
		return nil
	}
	return fmt.Errorf("未找到可用Redis配置")
}

func Client() *gredis.Redis {
	return client
}

func StockKey(activityId, goodsId uint32) string {
	return fmt.Sprintf("flash_sale:stock:%d:%d", activityId, goodsId)
}

func PurchaseKey(activityId, goodsId, userId uint32) string {
	return fmt.Sprintf("flash_sale:purchase:%d:%d:%d", activityId, goodsId, userId)
}

func ResultKey(resultId string) string {
	return fmt.Sprintf("flash_sale:result:%s", resultId)
}

func SetStock(ctx context.Context, activityId, goodsId uint32, stock uint) error {
	_, err := client.Do(ctx, "SET", StockKey(activityId, goodsId), stock)
	return err
}

func GetStock(ctx context.Context, activityId, goodsId uint32) (int, error) {
	value, err := client.Do(ctx, "GET", StockKey(activityId, goodsId))
	if err != nil {
		return 0, err
	}
	if value.IsNil() {
		return 0, fmt.Errorf("库存缓存不存在")
	}
	return value.Int(), nil
}

func ReduceStock(ctx context.Context, activityId, goodsId uint32, count uint32) error {
	script := `
local stock = redis.call("GET", KEYS[1])
if not stock then
  return -1
end
stock = tonumber(stock)
local count = tonumber(ARGV[1])
if stock < count then
  return 0
end
redis.call("DECRBY", KEYS[1], count)
return 1
`
	result, err := client.Do(ctx, "EVAL", script, 1, StockKey(activityId, goodsId), count)
	if err != nil {
		return err
	}
	switch result.Int() {
	case 1:
		return nil
	case 0:
		return fmt.Errorf("库存不足")
	default:
		return fmt.Errorf("库存缓存不存在")
	}
}

func IncreaseStock(ctx context.Context, activityId, goodsId uint32, count uint32) {
	if client == nil {
		return
	}
	if _, err := client.Do(ctx, "INCRBY", StockKey(activityId, goodsId), count); err != nil {
		g.Log().Errorf(ctx, "补偿Redis库存失败: %v", err)
	}
}

func CheckLimit(ctx context.Context, key string, limit int, ttl time.Duration) error {
	script := `
local current = redis.call("INCR", KEYS[1])
if current == 1 then
  redis.call("EXPIRE", KEYS[1], ARGV[2])
end
if current > tonumber(ARGV[1]) then
  return 0
end
return 1
`
	result, err := client.Do(ctx, "EVAL", script, 1, key, limit, int(ttl.Seconds()))
	if err != nil {
		return err
	}
	if result.Int() != 1 {
		return fmt.Errorf("请求过于频繁，请稍后再试")
	}
	return nil
}

func Exists(ctx context.Context, key string) (bool, error) {
	result, err := client.Do(ctx, "EXISTS", key)
	if err != nil {
		return false, err
	}
	return result.Int() > 0, nil
}

func SetPurchase(ctx context.Context, activityId, goodsId, userId uint32, ttl time.Duration) error {
	_, err := client.Do(ctx, "SET", PurchaseKey(activityId, goodsId, userId), 1, "EX", int(ttl.Seconds()))
	return err
}

func SetJSON(ctx context.Context, key string, value any, ttl time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if ttl > 0 {
		_, err = client.Do(ctx, "SET", key, string(bytes), "EX", int(ttl.Seconds()))
	} else {
		_, err = client.Do(ctx, "SET", key, string(bytes))
	}
	return err
}

func GetJSON(ctx context.Context, key string, value any) (bool, error) {
	result, err := client.Do(ctx, "GET", key)
	if err != nil {
		return false, err
	}
	if result.IsNil() {
		return false, nil
	}
	if err = json.Unmarshal([]byte(result.String()), value); err != nil {
		return false, err
	}
	return true, nil
}
