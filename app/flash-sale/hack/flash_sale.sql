CREATE DATABASE IF NOT EXISTS `flash_sale` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `flash_sale`;

CREATE TABLE IF NOT EXISTS `flash_sale_goods` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '秒杀商品ID',
  `goods_id` bigint unsigned NOT NULL COMMENT '商品ID',
  `activity_id` bigint unsigned NOT NULL COMMENT '活动ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '秒杀标题',
  `description` varchar(500) NOT NULL DEFAULT '' COMMENT '秒杀描述',
  `original_price` bigint unsigned NOT NULL DEFAULT 0 COMMENT '原价，单位分',
  `sale_price` bigint unsigned NOT NULL DEFAULT 0 COMMENT '秒杀价，单位分',
  `total_stock` int unsigned NOT NULL DEFAULT 0 COMMENT '总库存',
  `available_stock` int unsigned NOT NULL DEFAULT 0 COMMENT '可用库存',
  `start_time` datetime NOT NULL COMMENT '开始时间',
  `end_time` datetime NOT NULL COMMENT '结束时间',
  `status` tinyint unsigned NOT NULL DEFAULT 1 COMMENT '状态 1启用 2禁用 3结束',
  `image_url` varchar(500) NOT NULL DEFAULT '' COMMENT '商品图片URL',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_activity_goods` (`activity_id`,`goods_id`),
  KEY `idx_status_time` (`status`,`start_time`,`end_time`),
  KEY `idx_goods_id` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='秒杀商品表';

CREATE TABLE IF NOT EXISTS `flash_sale_order` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '秒杀订单ID',
  `order_no` varchar(64) NOT NULL DEFAULT '' COMMENT '秒杀订单号',
  `goods_id` bigint unsigned NOT NULL COMMENT '商品ID',
  `activity_id` bigint unsigned NOT NULL COMMENT '活动ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `count` int unsigned NOT NULL DEFAULT 1 COMMENT '购买数量',
  `amount` bigint unsigned NOT NULL DEFAULT 0 COMMENT '实付金额，单位分',
  `status` tinyint unsigned NOT NULL DEFAULT 1 COMMENT '状态 1成功 2失败 3取消',
  `result_id` varchar(64) NOT NULL DEFAULT '' COMMENT '秒杀结果ID',
  `message` varchar(255) NOT NULL DEFAULT '' COMMENT '处理消息',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_order_no` (`order_no`),
  UNIQUE KEY `uk_activity_goods_user` (`activity_id`,`goods_id`,`user_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_result_id` (`result_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='秒杀订单表';

CREATE TABLE IF NOT EXISTS `flash_sale_result` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '秒杀结果ID',
  `result_id` varchar(64) NOT NULL DEFAULT '' COMMENT '业务结果ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `goods_id` bigint unsigned NOT NULL COMMENT '商品ID',
  `activity_id` bigint unsigned NOT NULL COMMENT '活动ID',
  `order_no` varchar(64) NOT NULL DEFAULT '' COMMENT '秒杀订单号',
  `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态 0处理中 1成功 2失败',
  `message` varchar(255) NOT NULL DEFAULT '' COMMENT '处理消息',
  `pay_amount` bigint unsigned NOT NULL DEFAULT 0 COMMENT '支付金额，单位分',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_result_id` (`result_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_activity_goods_user` (`activity_id`,`goods_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='秒杀结果表';
