package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
)

const orderTimeoutType = "order_timeout"

// OrderGoodsInfo is the goods payload used by order domain events.
type OrderGoodsInfo struct {
	GoodsId int `json:"goods_id"`
	Count   int `json:"count"`
}

// OrderCreatedEvent is published after an order is created.
type OrderCreatedEvent struct {
	UserId    uint32            `json:"user_id"`
	OrderId   uint32            `json:"order_id"`
	GoodsIds  []uint32          `json:"goods_ids"`
	GoodsInfo []*OrderGoodsInfo `json:"goods_info"`
}

type couponConfirmEvent struct {
	OrderId  int `json:"order_id"`
	UserID   int `json:"user_id"`
	CouponId int `json:"coupon_id"`
}

type orderTimeoutEvent struct {
	OrderId   int    `json:"order_id"`
	Type      string `json:"type"`
	TimeStamp string `json:"timestamp"`
}

// PublishOrderCreatedEvent publishes an order-created event if the exchange is configured.
func PublishOrderCreatedEvent(event OrderCreatedEvent) {
	ctx := context.Background()
	publish(ctx, "rabbitmq.default.exchange.orderExchange", "rabbitmq.default.routingKey.orderCreated", event, 0)
}

// PublishCouponConfirmEvent asks goods service to confirm coupon usage.
func PublishCouponConfirmEvent(orderId int32, userID int32, couponId int32) {
	ctx := context.Background()
	event := couponConfirmEvent{
		OrderId:  int(orderId),
		UserID:   int(userID),
		CouponId: int(couponId),
	}
	publish(ctx, "rabbitmq.default.exchange.couponConfirmExchange", "rabbitmq.default.routingKey.couponConfirm", event, 0)
}

// PublishOrderTimeoutEvent publishes delayed order-timeout event.
func PublishOrderTimeoutEvent(orderId int, delayMs int) {
	ctx := context.Background()
	event := orderTimeoutEvent{
		OrderId:   orderId,
		Type:      orderTimeoutType,
		TimeStamp: time.Now().Format(time.RFC3339),
	}
	publish(ctx, "rabbitmq.default.exchange.orderDelayExchange", "rabbitmq.default.routingKey.orderTimeout", event, delayMs)
}

func publish(ctx context.Context, exchangeKey string, routingKeyKey string, event interface{}, delayMs int) {
	exchange := g.Cfg().MustGet(ctx, exchangeKey).String()
	routingKey := g.Cfg().MustGet(ctx, routingKeyKey).String()
	if exchange == "" || routingKey == "" {
		g.Log().Warningf(ctx, "RabbitMQ配置缺失，跳过事件发布: exchangeKey=%s routingKeyKey=%s", exchangeKey, routingKeyKey)
		return
	}

	client, err := GetOrderRabbitMQClient(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "连接RabbitMQ失败，跳过事件发布: %v", err)
		return
	}

	body, err := json.Marshal(event)
	if err != nil {
		g.Log().Errorf(ctx, "序列化RabbitMQ事件失败: %v", err)
		return
	}

	publishing := amqp.Publishing{
		ContentType:  "application/json",
		Body:         body,
		DeliveryMode: amqp.Persistent,
	}
	if delayMs > 0 {
		publishing.Headers = amqp.Table{"x-delay": delayMs}
	}

	exchangeType := "topic"
	exchangeArgs := amqp.Table(nil)
	if delayMs > 0 {
		exchangeType = "x-delayed-message"
		exchangeArgs = amqp.Table{"x-delayed-type": "direct"}
	}
	if err = client.channel.ExchangeDeclare(exchange, exchangeType, true, false, false, false, exchangeArgs); err != nil {
		g.Log().Errorf(ctx, "声明RabbitMQ交换机失败: %v", err)
		return
	}

	if err = client.channel.Publish(exchange, routingKey, false, false, publishing); err != nil {
		g.Log().Errorf(ctx, "发布RabbitMQ事件失败: %v", err)
	}
}
