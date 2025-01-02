
# sqlite3

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for pre_user
-- ----------------------------
DROP TABLE IF EXISTS "pre_user";
CREATE TABLE "pre_user" (
    "uid" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "email" TEXT DEFAULT '',
    "username" TEXT DEFAULT '',
    "password" TEXT DEFAULT '',
    `salt` TEXT DEFAULT '',
    "token" TEXT NOT NULL DEFAULT '',
    "avatar" TEXT NOT NULL DEFAULT '',
    "is_deleted" INTEGER NOT NULL DEFAULT 0,
    "updated_at" TIMESTAMP DEFAULT NULL,
    "created_at" TIMESTAMP DEFAULT NULL,
    "deleted_at" TIMESTAMP DEFAULT NULL,
    UNIQUE ("email" ASC)
);
CREATE UNIQUE INDEX "email" ON "pre_user" ("email" ASC);

DROP TABLE IF EXISTS "pre_cloud_config";
CREATE TABLE "pre_cloud_config" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "uid" INTEGER NOT NULL DEFAULT 0,
    "bucket_name" TEXT DEFAULT '',
    "account_id" TEXT DEFAULT '',
    "access_key_id" TEXT DEFAULT '',
    "access_key_secret" TEXT DEFAULT '',
    "custom_path" TEXT DEFAULT '',
    "is_deleted" INTEGER NOT NULL DEFAULT 0,
    "updated_at" TIMESTAMP DEFAULT NULL,
    "created_at" TIMESTAMP DEFAULT NULL,
    "deleted_at" TIMESTAMP DEFAULT NULL
);
CREATE INDEX "uid" ON "pre_cloud_config" ("uid" ASC);




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


