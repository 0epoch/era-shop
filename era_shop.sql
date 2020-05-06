CREATE DATABASE IF NOT EXISTS `era_shop` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `era_shop`;

-- 导出  表 era_shop.era_accounts 结构
CREATE TABLE IF NOT EXISTS `era_accounts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL DEFAULT '',
  `password` varchar(64) NOT NULL DEFAULT '',
  `phone` varchar(32) NOT NULL DEFAULT '',
  `email` varchar(64) NOT NULL DEFAULT '',
  `last_login` datetime DEFAULT NULL COMMENT '最后登录时间',
  `last_ip` char(50) DEFAULT '' COMMENT '最后登录IP',
  `created_ip` char(50) DEFAULT '' COMMENT '注册IP',
  `status` tinyint(4) DEFAULT '0',
  `created_at` datetime DEFAULT NULL COMMENT '注册时间',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `era_accounts_phone_IDX` (`phone`) USING BTREE,
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='账户';

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_admins 结构
CREATE TABLE IF NOT EXISTS `era_admins` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(64) NOT NULL DEFAULT '',
  `phone` varchar(32) NOT NULL DEFAULT '',
  `avatar` varchar(64) NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  KEY `account_id` (`account_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='管理员';

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_admin_roles 结构
CREATE TABLE IF NOT EXISTS `era_admin_roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT '0',
  `admin_id` int(11) NOT NULL DEFAULT '0',
  `roles` varchar(155) NOT NULL DEFAULT '' COMMENT '角色ID，多个用逗号隔开',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='管理员-角色关联表';

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_attrs 结构
CREATE TABLE IF NOT EXISTS `era_attrs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '',
  `desc` varchar(255) DEFAULT '',
  `status` tinyint(4) DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_attr_values 结构
CREATE TABLE IF NOT EXISTS `era_attr_values` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `attr_id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(50) NOT NULL DEFAULT '',
  `desc` varchar(255) NOT NULL DEFAULT '',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_brands 结构
CREATE TABLE IF NOT EXISTS `era_brands` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `logo` varchar(64) NOT NULL DEFAULT '' COMMENT '品牌logo uri',
  `status` tinyint(4) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='品牌';

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_carts 结构
CREATE TABLE IF NOT EXISTS `era_carts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT '0',
  `spu_id` int(11) NOT NULL DEFAULT '0',
  `sku_id` int(11) NOT NULL DEFAULT '0',
  `quantity` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '单价',
  `sku_snapshot` json DEFAULT NULL COMMENT '商品快照',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `spu_id` (`spu_id`,`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_categories 结构
CREATE TABLE IF NOT EXISTS `era_categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '',
  `brands` varchar(255) NOT NULL DEFAULT '' COMMENT '多个品牌ID用逗号隔开',
  `pid` int(11) NOT NULL,
  `image` varchar(64) NOT NULL DEFAULT '' COMMENT '图标',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `path` varchar(255) NOT NULL DEFAULT '' COMMENT 'pid-id',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='分类';

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_order 结构
CREATE TABLE IF NOT EXISTS `era_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_no` varchar(64) NOT NULL DEFAULT '',
  `account_id` int(11) NOT NULL DEFAULT '0',
  `products_num` int(11) NOT NULL DEFAULT '0' COMMENT '商品数量',
  `total` int(11) NOT NULL DEFAULT '0' COMMENT '实际总价',
  `pay_total` int(11) NOT NULL DEFAULT '0' COMMENT '支付总价',
  `revises_total` int(11) NOT NULL DEFAULT '0' COMMENT '调整金额',
  `status` enum('checked','created','cancelled','finished') NOT NULL DEFAULT 'checked' COMMENT '订单状态',
  `payment_status` enum('checked','await','cancelled','paid','part_refunded','refunded') NOT NULL DEFAULT 'checked' COMMENT '支付状态',
  `shipment_status` enum('checked','ready','cancelled','part_shipped','shipped') NOT NULL DEFAULT 'checked' COMMENT '物流状态',
  `user_ip` varchar(64) NOT NULL DEFAULT '',
  `paid_at` datetime DEFAULT NULL COMMENT '支付时间',
  `shipped_at` datetime DEFAULT NULL COMMENT '发货时间',
  `confirmed_at` datetime DEFAULT NULL COMMENT '确实收货时间',
  `reviewed_at` datetime DEFAULT NULL COMMENT '评论时间',
  `finished_at` datetime DEFAULT NULL COMMENT '订单完成时间',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_no` (`order_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_order_item 结构
CREATE TABLE IF NOT EXISTS `era_order_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_no` varchar(64) NOT NULL DEFAULT '',
  `spu_id` int(11) NOT NULL DEFAULT '0',
  `sku_id` int(11) NOT NULL DEFAULT '0',
  `quantity` mediumint(9) NOT NULL DEFAULT '0',
  `revises_total` int(11) NOT NULL DEFAULT '0' COMMENT '调整金额',
  `total` int(11) NOT NULL DEFAULT '0' COMMENT '订单项总价',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '单价',
  `rest` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `order_no` (`order_no`),
  KEY `spu_id_sku_id` (`spu_id`,`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单项';

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_permissions 结构
CREATE TABLE IF NOT EXISTS `era_permissions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT '0',
  `name` varchar(64) NOT NULL DEFAULT '',
  `uri` varchar(64) NOT NULL DEFAULT '' COMMENT '路由',
  `is_menu` tinyint(4) NOT NULL DEFAULT '0',
  `method` char(8) NOT NULL DEFAULT '' COMMENT '请求方式',
  `icon` varchar(50) NOT NULL DEFAULT '0',
  `order` smallint(6) NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='权限';

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_revises 结构
CREATE TABLE IF NOT EXISTS `era_revises` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_no` varchar(50) NOT NULL DEFAULT '',
  `order_item_id` int(11) NOT NULL DEFAULT '0',
  `type` char(50) NOT NULL DEFAULT '',
  `amount` int(11) NOT NULL DEFAULT '0' COMMENT '调整金额',
  `index_key` varchar(50) NOT NULL DEFAULT '' COMMENT '调整类型的主键',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '说明',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='价格调整记录';

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_roles 结构
CREATE TABLE IF NOT EXISTS `era_roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL,
  `permissions` varchar(255) NOT NULL DEFAULT '' COMMENT '权限ID，多个用逗号隔开',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='角色';

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_skus 结构
CREATE TABLE IF NOT EXISTS `era_skus` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `spu_id` int(11) NOT NULL,
  `attrs` varchar(50) NOT NULL DEFAULT '' COMMENT '属性值ID，多个逗号隔开',
  `stock` int(11) NOT NULL DEFAULT '0',
  `banner_img` varchar(50) NOT NULL DEFAULT '',
  `main_img` varchar(50) NOT NULL DEFAULT '',
  `price` int(11) NOT NULL DEFAULT '0',
  `market_price` int(11) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_spus 结构
CREATE TABLE IF NOT EXISTS `era_spus` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `brand_id` int(11) NOT NULL DEFAULT '0',
  `category_id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(64) NOT NULL DEFAULT '',
  `unit` char(8) NOT NULL DEFAULT '',
  `banner_img` varchar(64) NOT NULL DEFAULT '',
  `main_img` varchar(64) NOT NULL DEFAULT '',
  `price` int(11) NOT NULL DEFAULT '0',
  `market_price` int(11) NOT NULL DEFAULT '0',
  `desc` varchar(255) NOT NULL DEFAULT '',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `brand_id` (`brand_id`),
  KEY `cartegory_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_spu_sku_attr 结构
CREATE TABLE IF NOT EXISTS `era_spu_sku_attr` (
  `id` int(11) NOT NULL,
  `spu_id` int(11) NOT NULL DEFAULT '0',
  `sku_id` int(11) NOT NULL DEFAULT '0',
  `attr_id` int(11) NOT NULL DEFAULT '0',
  `attr_name` varchar(50) NOT NULL DEFAULT '',
  `attr_value_id` int(11) NOT NULL DEFAULT '0',
  `attr_value_name` varchar(50) NOT NULL DEFAULT '',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 数据导出被取消选择。

-- 导出  表 era_shop.era_users 结构
CREATE TABLE IF NOT EXISTS `era_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT '0',
  `nickname` varchar(64) NOT NULL DEFAULT '',
  `bio` varchar(50) NOT NULL DEFAULT '',
  `avatar` varchar(64) DEFAULT '',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='用户';

