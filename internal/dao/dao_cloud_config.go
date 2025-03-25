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
	ID              int64      `gorm:"column:id;primary_key;auto_increment" json:"id" form:"id"`                                            //
	UID             int64      `gorm:"column:uid;index;default:0" json:"uid" form:"uid"`                                                    //
	Type            string     `gorm:"column:type;default:''" json:"type" form:"type"`                                                      //
	BucketName      string     `gorm:"column:bucket_name;default:''" json:"bucketName" form:"bucketName"`                                   //
	Endpoint        string     `gorm:"column:endpoint;default:''" json:"endpoint" form:"endpoint"`                                          //
	Region          string     `gorm:"column:region;default:''" json:"region" form:"region"`                                                //
	AccountID       string     `gorm:"column:account_id;default:''" json:"accountId" form:"accountId"`                                      //
	AccessKeyID     string     `gorm:"column:access_key_id;default:''" json:"accessKeyId" form:"accessKeyId"`                               //
	AccessKeySecret string     `gorm:"column:access_key_secret;default:''" json:"accessKeySecret" form:"accessKeySecret"`                   //
	CustomPath      string     `gorm:"column:custom_path;default:''" json:"customPath" form:"customPath"`                                   //
	AccessURLPrefix string     `gorm:"column:access_url_prefix;default:''" json:"accessUrlPrefix" form:"accessUrlPrefix"`                   //
	IsEnabled       int64      `gorm:"column:is_enabled;default:1" json:"isEnabled" form:"isEnabled"`                                       //
	IsDeleted       int64      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                                       //
	UpdatedAt       timex.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime:false;default:NULL" json:"updatedAt" form:"updatedAt"` //
	CreatedAt       timex.Time `gorm:"column:created_at;type:datetime;autoUpdateTime:false;default:NULL" json:"createdAt" form:"createdAt"` //
	DeletedAt       timex.Time `gorm:"column:deleted_at;type:datetime;autoUpdateTime:false;default:NULL" json:"deletedAt" form:"deletedAt"` //
}

type CloudConfigSet struct {
	ID              int64  `gorm:"column:id;primary_key;auto_increment" json:"id" form:"id"`                          //
	Type            string `gorm:"column:type;default:''" json:"type" form:"type"`                                    //
	BucketName      string `gorm:"column:bucket_name;default:''" json:"bucketName" form:"bucketName"`                 //
	Endpoint        string `gorm:"column:endpoint;default:''" json:"endpoint" form:"endpoint"`                        //
	Region          string `gorm:"column:region;default:''" json:"region" form:"region"`                              //
	AccountID       string `gorm:"column:account_id;default:''" json:"accountId" form:"accountId"`                    //
	AccessKeyID     string `gorm:"column:access_key_id;default:''" json:"accessKeyId" form:"accessKeyId"`             //
	AccessKeySecret string `gorm:"column:access_key_secret;default:''" json:"accessKeySecret" form:"accessKeySecret"` //
	CustomPath      string `gorm:"column:custom_path;default:''" json:"customPath" form:"customPath"`                 //
	AccessURLPrefix string `gorm:"column:access_url_prefix;default:''" json:"accessUrlPrefix" form:"accessUrlPrefix"` //
	IsEnabled       int64  `gorm:"column:is_enabled;default:1" json:"isEnabled" form:"isEnabled"`                     //
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
	err := u.WithContext(d.ctx).Create(m)
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
