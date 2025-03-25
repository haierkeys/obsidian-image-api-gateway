
# sqlite3

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for pre_user
-- ----------------------------
DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
    "uid" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "email" text DEFAULT '',
    "username" text DEFAULT '',
    "password" text DEFAULT '',
    "salt" text DEFAULT '',
    "token" text NOT NULL DEFAULT '',
    "avatar" text NOT NULL DEFAULT '',
    "is_deleted" integer NOT NULL DEFAULT 0,
    "updated_at" datetime DEFAULT NULL,
    "created_at" datetime DEFAULT NULL,
    "deleted_at" datetime DEFAULT NULL,
    UNIQUE ("email" ASC)
);
CREATE UNIQUE INDEX "idx_user_email" ON "user" ("email" ASC);

DROP TABLE IF EXISTS "cloud_config";
CREATE TABLE "cloud_config" (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "uid" integer NOT NULL DEFAULT 0,
    "type" text DEFAULT '',
    "endpoint" text DEFAULT '',
    "region" text DEFAULT '',
    "account_id" text DEFAULT '',
    "bucket_name" text DEFAULT '',
    "access_key_id" text DEFAULT '',
    "access_key_secret" text DEFAULT '',
    "custom_path" text DEFAULT '',
    "access_url_prefix" text DEFAULT '',
    "user" text DEFAULT '',
    "password" text DEFAULT '',
    "path" text DEFAULT '',

    "is_enabled" integer NOT NULL DEFAULT 1,
    "is_deleted" integer NOT NULL DEFAULT 0,
    "updated_at" datetime DEFAULT NULL,
    "created_at" datetime DEFAULT NULL,
    "deleted_at" datetime DEFAULT NULL
);
CREATE INDEX "idx_cloud_config_uid" ON "cloud_config" ("uid" ASC);


PRAGMA foreign_keys = true;


## mysql
DROP TABLE IF EXISTS `pre_user`;
CREATE TABLE `pre_user`  (
 `uid` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
 `email` char(255)  NOT NULL DEFAULT '' COMMENT '邮箱地址',
 `username` char(255)  NOT NULL DEFAULT '' COMMENT '用户名',
 `password` char(32)  NOT NULL DEFAULT '' COMMENT '密码',
 `salt` char(24)  NOT NULL DEFAULT '' COMMENT '密码混淆码',
 `token` char(255)  NOT NULL DEFAULT '' COMMENT '用户授权令牌',
 `avatar` char(255)  NOT NULL DEFAULT '' COMMENT '用户头像路径',
 `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
 `updated_at` datetime(0) NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
 `created_at` datetime(0) NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
 `deleted_at` datetime(0) NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '标记删除时间',
 PRIMARY KEY (`uid`) ,
 UNIQUE INDEX `email`(`email`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4  COMMENT = '用户表';

DROP TABLE IF EXISTS "pre_cloud_config";

CREATE TABLE `pre_cloud_config` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '配置id',
    `uid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `type` char(255) NOT NULL DEFAULT '' COMMENT '类型',
    `bucket_name` char(255) NOT NULL DEFAULT '' COMMENT '存储桶名称',
    `account_id` char(255) NOT NULL DEFAULT '' COMMENT '账户id',
    `access_key_id` char(255) NOT NULL DEFAULT '' COMMENT '访问密钥id',
    `access_key_secret` char(255) NOT NULL DEFAULT '' COMMENT '访问密钥',
    `custom_path` char(255) NOT NULL DEFAULT '' COMMENT '自定义路径',
    `access_url_prefix` char(255) NOT NULL DEFAULT '' COMMENT '访问地址前缀',
    `is_enabled` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否启用',
    `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
    `updated_at` datetime(0) NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
    `created_at` datetime(0) NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
    `deleted_at` datetime(0) NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '标记删除时间',
    PRIMARY KEY (`id`),
    INDEX `uid`(`uid`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COMMENT = '云配置表';
