#### user 
用户表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | uid | 用户UID | bigint(20) unsigned | PRI | NO | auto_increment |  |
| 2 | nickname | 用户昵称 | char(255) |  | NO |  | '' |
| 3 | avatar | 头像 | char(255) |  | NO |  | '' |
| 4 | mobile | 电话号码 | char(255) | MUL | NO |  | '' |
| 5 | name | 真实姓名 | char(255) |  | NO |  | '' |
| 6 | idcard | 身份证号 | char(255) |  | NO |  | '' |
| 7 | is_validate | 是否身份验证 | tinyint(1) unsigned |  | NO |  | 0 |
| 8 | openid | openid | varchar(255) | MUL | NO |  | '' |
| 9 | unionid | unionid | varchar(255) | MUL | NO |  | '' |
| 10 | gender | 性别 | varchar(255) |  | NO |  | '' |
| 11 | language | 语言 | varchar(255) |  | NO |  | '' |
| 12 | city | 城市 | varchar(255) |  | NO |  | '' |
| 13 | province | 省份 | varchar(255) |  | NO |  | '' |
| 14 | country | 国家 | varchar(255) |  | NO |  | '' |
| 15 | avatar_url | 微信的头像 | varchar(255) |  | NO |  | '' |
| 16 | session_key | session_key | varchar(255) |  | NO |  | '' |
| 17 | token | 用户TOKEN | char(255) |  | NO |  | '' |
| 18 | weixin_token | 用户微信TOKEN | char(255) |  | NO |  | '' |
| 19 | change_name_num | 用户昵称剩余修改次数 | tinyint(1) unsigned |  | NO |  | 0 |
| 20 | hannels_id | 渠道ID | bigint(20) | MUL | NO |  | 0 |
| 21 | is_deleted | 是否删除 | tinyint(1) unsigned |  | NO |  | 0 |
| 22 | updated_at | 更新时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 23 | created_at | 创建时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 24 | deleted_at | 标记删除时间 | datetime |  | NO |  | '0000-00-00 00:00:00' |
| 25 | app_id | 注册来源 应用的ID | bigint(20) unsigned | MUL | NO |  | 0 |
