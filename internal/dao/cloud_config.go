package dao

import (
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model/main_gen/cloud_config_repo"

	"gorm.io/gorm"
)

type CloudConfigDao struct {
	db *gorm.DB
}

func NewCloudConfigDao(db *gorm.DB) *CloudConfigDao {
	return &CloudConfigDao{
		db: db,
	}
}

// Create 创建云存储配置
func (d *CloudConfigDao) Create(config *cloud_config_repo.CloudConfig) (int64, error) {
	return config.Create()
}

// Update 更新云存储配置
func (d *CloudConfigDao) Update(config *cloud_config_repo.CloudConfig) error {
	return config.Save()
}

// Delete 删除云存储配置
func (d *CloudConfigDao) Delete(uid int64, id int64) error {
	return cloud_config_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereId(model.Eq, id).
		Delete()
}

// GetByID 根据ID获取配置
func (d *CloudConfigDao) GetByID(id int64) (*cloud_config_repo.CloudConfig, error) {
	return cloud_config_repo.NewQueryBuilder().
		WhereId(model.Eq, id).
		QueryOne()
}

// GetByUID 获取用户的所有配置
func (d *CloudConfigDao) GetByUID(uid int64) ([]*cloud_config_repo.CloudConfig, error) {
	return cloud_config_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		OrderByCreatedAt(false).
		QueryAll()
}

// GetByType 获取指定类型的配置
func (d *CloudConfigDao) GetByType(uid int64, typeStr string) ([]*cloud_config_repo.CloudConfig, error) {
	return cloud_config_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereType(model.Eq, typeStr).
		OrderByCreatedAt(false).
		QueryAll()
}

// GetByUIDAndType 获取用户指定类型的单个配置
func (d *CloudConfigDao) GetByUIDAndType(uid int64, typeStr string) (*cloud_config_repo.CloudConfig, error) {
	return cloud_config_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereType(model.Eq, typeStr).
		QueryOne()
}
