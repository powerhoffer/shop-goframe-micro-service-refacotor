-- 微信支付/退款所需字段补丁。
-- 如果表结构已经包含这些字段，可以忽略对应 ALTER 报错或手动确认后跳过。

ALTER TABLE `order_info`
  ADD COLUMN `transaction_id` varchar(100) NOT NULL DEFAULT '' COMMENT '第三方支付交易号' AFTER `status`;

ALTER TABLE `refund_info`
  ADD COLUMN `refund_id` varchar(100) NOT NULL DEFAULT '' COMMENT '第三方退款编号' AFTER `goods_id`;

ALTER TABLE `refund_info`
  ADD COLUMN `refund_status` tinyint NOT NULL DEFAULT 0 COMMENT '退款状态 0未退款 1退款中 2退款成功 3退款失败' AFTER `status`;

ALTER TABLE `refund_info`
  ADD COLUMN `refund_amount` int NOT NULL DEFAULT 0 COMMENT '退款金额 单位分' AFTER `refund_status`;
