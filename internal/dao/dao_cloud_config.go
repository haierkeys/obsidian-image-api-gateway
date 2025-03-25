package dao

import (
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/query"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/app"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/convert"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/timex"
	"gorm.io/gorm"
)

type CloudConfig struct {
	ID              int64      `json:"id" form:"id"`                           // 主键ID
	UID             int64      `json:"uid" form:"uid"`                         // 用户ID，非空，索引
	Type            string     `json:"type" form:"type"`                       // 类型
	Endpoint        string     `json:"endpoint" form:"endpoint"`               // 终端点
	Region          string     `json:"region" form:"region"`                   // 区域
	AccountID       string     `json:"accountId" form:"accountId"`             // 账户ID
	BucketName      string     `json:"bucketName" form:"bucketName"`           // 桶名称
	AccessKeyID     string     `json:"accessKeyId" form:"accessKeyId"`         // 访问密钥ID
	AccessKeySecret string     `json:"accessKeySecret" form:"accessKeySecret"` // 访问密钥密文
	CustomPath      string     `json:"customPath" form:"customPath"`           // 自定义路径
	AccessURLPrefix string     `json:"accessUrlPrefix" form:"accessUrlPrefix"` // 访问URL前缀
	User            string     `json:"user" form:"user"`                       // 用户名
	Password        string     `json:"password" form:"password"`               // 密码
	IsEnabled       int64      `json:"isEnabled" form:"isEnabled"`             // 是否启用，非空，默认为1
	IsDeleted       int64      `json:"isDeleted" form:"isDeleted"`             // 是否删除，非空
	UpdatedAt       timex.Time `json:"updatedAt" form:"updatedAt"`             // 更新时间，自动更新时间戳
	CreatedAt       timex.Time `json:"createdAt" form:"createdAt"`             // 创建时间，自动创建时间戳
	DeletedAt       timex.Time `json:"deletedAt" form:"deletedAt"`             // 删除时间，默认为NULL
}

type CloudConfigSet struct {
	ID              int64  `json:"id" form:"id"`                           // 主键ID
	Type            string `json:"type" form:"type"`                       // 类型
	Endpoint        string `json:"endpoint" form:"endpoint"`               // 终端点
	Region          string `json:"region" form:"region"`                   // 区域
	AccountID       string `json:"accountId" form:"accountId"`             // 账户ID
	BucketName      string `json:"bucketName" form:"bucketName"`           // 桶名称
	AccessKeyID     string `json:"accessKeyId" form:"accessKeyId"`         // 访问密钥ID
	AccessKeySecret string `json:"accessKeySecret" form:"accessKeySecret"` // 访问密钥密文
	CustomPath      string `json:"customPath" form:"customPath"`           // 自定义路径
	AccessURLPrefix string `json:"accessUrlPrefix" form:"accessUrlPrefix"` // 访问URL前缀
	User            string `json:"user" form:"user"`                       // 用户名
	Password        string `json:"password" form:"password"`               // 密码
	IsEnabled       int64  `json:"isEnabled" form:"isEnabled"`             // 是否启用，非空，默认为1
}

func (d *Dao) cloudConfig() *query.Query {
	return d.Use(
		func(g *gorm.DB) {
			model.AutoMigrate(g, "CloudConfig")
		},
	)
}

// 创建云存储配置
func (d *Dao) Create(params *CloudConfigSet, uid int64) (int64, error) {

	u := d.cloudConfig().CloudConfig

	m := convert.StructAssign(params, &model.CloudConfig{}).(*model.CloudConfig)
	m.UID = uid
	err := u.WithContext(d.ctx).Create(m)

	//dump.P(m, uid)

	if err != nil {
		return 0, err
	}
	return m.ID, nil

}

// 更新云存储配置
func (d *Dao) Update(params *CloudConfigSet, id int64, uid int64) error {

	u := d.cloudConfig().CloudConfig

	m, err := u.WithContext(d.ctx).Where(
		u.ID.Eq(id),
		u.UID.Eq(uid),
		u.IsDeleted.Eq(0),
	).First()
	if err != nil {
		return err
	}

	convert.StructAssign(params, m)

	m.ID = id
	m.UID = uid

	err = u.WithContext(d.ctx).Where(u.ID.Eq(id)).Save(m)
	return err
}

// 启用云存储配置
func (d *Dao) Enable(id int64, uid int64) error {
	u := d.cloudConfig().CloudConfig

	_, err := u.WithContext(d.ctx).Where(
		u.ID.Eq(id),
		u.UID.Eq(uid),
		u.IsDeleted.Eq(0),
	).UpdateSimple(
		u.IsEnabled.Value(1),
		u.UpdatedAt.Value(timex.Now()),
	)
	return err
}

// 批量关闭云存储配置
func (d *Dao) DisableBatch(id int64, uid int64) error {

	u := d.cloudConfig().CloudConfig

	_, err := u.WithContext(d.ctx).Where(
		u.ID.Neq(id),
		u.UID.Eq(uid),
		u.IsDeleted.Eq(0),
	).UpdateSimple(
		u.IsEnabled.Value(0),
		u.UpdatedAt.Value(timex.Now()),
	)

	return err
}

func (d *Dao) CountListByUID(uid int64) (int64, error) {
	u := d.cloudConfig().CloudConfig

	return u.WithContext(d.ctx).Where(
		u.UID.Eq(uid),
		u.IsDeleted.Eq(0),
	).Count()
}

// 获取用户的云存储配置列表
func (d *Dao) GetListByUID(page int, pageSize int, uid int64) ([]*CloudConfig, error) {

	u := d.cloudConfig().CloudConfig

	modelList, err := u.WithContext(d.ctx).Where(
		u.UID.Eq(uid),
		u.IsDeleted.Eq(0),
	).Order(u.CreatedAt).
		Limit(pageSize).
		Offset(app.GetPageOffset(page, pageSize)).
		Find()

	if err != nil {
		return nil, err
	}

	var list []*CloudConfig
	for _, m := range modelList {
		list = append(list, convert.StructAssign(m, &CloudConfig{}).(*CloudConfig))
	}
	return list, nil
}

// 根据ID获取配置
func (d *Dao) GetEnableByUId(uid int64) (*CloudConfig, error) {

	u := d.cloudConfig().CloudConfig

	m, err := u.WithContext(d.ctx).Where(
		u.UID.Eq(uid),
		u.IsEnabled.Eq(1),
		u.IsDeleted.Eq(0),
	).First()

	if err != nil {
		return nil, err
	}
	return convert.StructAssign(m, &CloudConfig{}).(*CloudConfig), nil
}

// 根据ID获取配置
func (d *Dao) GetById(id int64, uid int64) (*CloudConfig, error) {

	u := d.cloudConfig().CloudConfig

	m, err := u.WithContext(d.ctx).Where(
		u.ID.Eq(id),
		u.UID.Eq(uid),
		u.IsDeleted.Eq(0),
	).First()

	if err != nil {
		return nil, err
	}
	return convert.StructAssign(m, &CloudConfig{}).(*CloudConfig), nil
}

// 删除配置
func (d *Dao) Delete(id int64, uid int64) error {
	u := d.cloudConfig().CloudConfig

	_, err := u.WithContext(d.ctx).Where(
		u.ID.Eq(id),
		u.UID.Eq(uid),
	).UpdateSimple(
		u.IsDeleted.Value(1),
		u.DeletedAt.Value(timex.Now()),
	)
	return err
}
