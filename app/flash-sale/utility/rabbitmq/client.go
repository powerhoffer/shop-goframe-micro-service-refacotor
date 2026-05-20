package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

type FlashSaleOrderMessage struct {
	ResultId   string `json:"result_id"`
	OrderNo    string `json:"order_no"`
	GoodsId    uint32 `json:"goods_id"`
	ActivityId uint32 `json:"activity_id"`
	UserId     uint32 `json:"user_id"`
	Count      uint32 `json:"count"`
	Amount     uint64 `json:"amount"`
}

var client *Client

func Init(ctx context.Context) error {
	address := g.Cfg().MustGet(ctx, "rabbitmq.default.address").String()
	if address == "" {
		return fmt.Errorf("RabbitMQ连接地址不能为空")
	}

	conn, err := amqp.Dial(address)
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return err
	}

	c := &Client{conn: conn, channel: ch}
	if err = c.initTopology(ctx); err != nil {
		_ = c.Close()
		return err
	}
	client = c
	g.Log().Info(ctx, "秒杀RabbitMQ初始化成功")
	return nil
}

func PublishFlashSaleOrder(ctx context.Context, msg FlashSaleOrderMessage) error {
	if client == nil {
		return fmt.Errorf("RabbitMQ未初始化")
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	exchange := g.Cfg().MustGet(ctx, "rabbitmq.default.exchange.flashSaleExchange").String()
	routingKey := g.Cfg().MustGet(ctx, "rabbitmq.default.routingKey.flashSaleOrder").String()
	return client.channel.PublishWithContext(ctx, exchange, routingKey, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
}

func (c *Client) initTopology(ctx context.Context) error {
	exchange := g.Cfg().MustGet(ctx, "rabbitmq.default.exchange.flashSaleExchange").String()
	queue := g.Cfg().MustGet(ctx, "rabbitmq.default.queue.flashSaleOrderQueue").String()
	routingKey := g.Cfg().MustGet(ctx, "rabbitmq.default.routingKey.flashSaleOrder").String()

	if exchange == "" || queue == "" || routingKey == "" {
		return fmt.Errorf("RabbitMQ秒杀交换机/队列/路由键配置不能为空")
	}
	if err := c.channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		return err
	}
	if _, err := c.channel.QueueDeclare(queue, true, false, false, false, nil); err != nil {
		return err
	}
	return c.channel.QueueBind(queue, routingKey, exchange, false, nil)
}

func (c *Client) Close() error {
	if c.channel != nil {
		_ = c.channel.Close()
	}
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
