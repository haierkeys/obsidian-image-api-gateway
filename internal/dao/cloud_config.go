package dao

import (
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model/main_gen/cloud_config_repo"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/app"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/convert"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/timef"
)

type CloudConfig struct {
	Id              int64      `gorm:"column:id;primary_key;auto_increment" json:"id" form:"id"`                                            //
	Uid             int64      `gorm:"column:uid;index;default:0" json:"uid" form:"uid"`                                                    //
	Type            string     `gorm:"column:type;default:''" json:"type" form:"type"`                                                      //
	BucketName      string     `gorm:"column:bucket_name;default:''" json:"bucketName" form:"bucketName"`                                   //
	AccountId       string     `gorm:"column:account_id;default:''" json:"accountId" form:"accountId"`                                      //
	AccessKeyId     string     `gorm:"column:access_key_id;default:''" json:"accessKeyId" form:"accessKeyId"`                               //
	AccessKeySecret string     `gorm:"column:access_key_secret;default:''" json:"accessKeySecret" form:"accessKeySecret"`                   //
	CustomPath      string     `gorm:"column:custom_path;default:''" json:"customPath" form:"customPath"`                                   //
	IsDeleted       int64      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                                       //
	UpdatedAt       timef.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime:false;default:NULL" json:"updatedAt" form:"updatedAt"` //
	CreatedAt       timef.Time `gorm:"column:created_at;type:datetime;autoUpdateTime:false;default:NULL" json:"createdAt" form:"createdAt"` //
	DeletedAt       timef.Time `gorm:"column:deleted_at;type:datetime;autoUpdateTime:false;default:NULL" json:"deletedAt" form:"deletedAt"` //
}

type CloudConfigSet struct {
	Id              int64  `gorm:"column:id;primary_key;auto_increment" json:"id" form:"id"`                          //
	Type            string `gorm:"column:type;default:''" json:"type" form:"type"`                                    //
	BucketName      string `gorm:"column:bucket_name;default:''" json:"bucketName" form:"bucketName"`                 //
	AccountId       string `gorm:"column:account_id;default:''" json:"accountId" form:"accountId"`                    //
	AccessKeyId     string `gorm:"column:access_key_id;default:''" json:"accessKeyId" form:"accessKeyId"`             //
	AccessKeySecret string `gorm:"column:access_key_secret;default:''" json:"accessKeySecret" form:"accessKeySecret"` //
	CustomPath      string `gorm:"column:custom_path;default:''" json:"customPath" form:"customPath"`                 //
}

// 创建云存储配置
func (d *Dao) Create(params *CloudConfigSet, uid int64) (int64, error) {

	m := &cloud_config_repo.CloudConfig{}
	convert.StructAssign(params, m)
	m.Uid = uid

	id, err := m.Create()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新云存储配置
func (d *Dao) Update(params *CloudConfigSet, uid int64, id int64) error {

	m, err := cloud_config_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereId(model.Eq, id).
		WhereIsDeleted(model.Eq, 0).
		First()
	if err != nil {
		return err
	}
	convert.StructAssign(params, m)
	m.Uid = uid
	m.Id = id
	err = m.Save()

	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) CountListByUid(uid int64) (int64, error) {
	return cloud_config_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereIsDeleted(model.Eq, 0).
		Count()
}

// 获取用户的云存储配置列表
func (d *Dao) GetListByUid(uid int64, page int, pageSize int) ([]*CloudConfig, error) {

	modelList, err := cloud_config_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereIsDeleted(model.Eq, 0).
		OrderByCreatedAt(false).
		Offset(app.GetPageOffset(page, pageSize)).
		Limit(pageSize).
		Get()

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
func (d *Dao) GetById(id int64, uid int64) (*CloudConfig, error) {

	m, err := cloud_config_repo.NewQueryBuilder().
		WhereId(model.Eq, id).
		WhereUid(model.Eq, uid).
		WhereIsDeleted(model.Eq, 0).
		First()
	if err != nil {
		return nil, err
	}
	return convert.StructAssign(m, &CloudConfig{}).(*CloudConfig), nil
}

// 删除配置
func (d *Dao) Delete(id int64, uid int64) error {
	return cloud_config_repo.NewQueryBuilder().
		WhereId(model.Eq, id).
		WhereUid(model.Eq, uid).
		Updates(map[string]interface{}{
			"is_deleted": 1,
			"deleted_at": timef.Now(),
		})
}
