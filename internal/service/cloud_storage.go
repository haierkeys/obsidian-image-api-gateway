package service

import (
	"errors"

	"github.com/haierkeys/obsidian-image-api-gateway/internal/dao"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model/main_gen/cloud_config_repo"
)

type CloudStorageService struct {
	cloudConfigDao *dao.CloudConfigDao
}

func NewCloudStorageService(cloudConfigDao *dao.CloudConfigDao) *CloudStorageService {
	return &CloudStorageService{
		cloudConfigDao: cloudConfigDao,
	}
}

type CreateCloudStorageRequest struct {
	Uid             int64  `json:"uid" binding:"required"`
	Type            string `json:"type" binding:"required"`
	BucketName      string `json:"bucket_name" binding:"required"`
	AccountId       string `json:"account_id"`
	AccessKeyId     string `json:"access_key_id" binding:"required"`
	AccessKeySecret string `json:"access_key_secret" binding:"required"`
	CustomPath      string `json:"custom_path"`
}

// Create 创建云存储配置
func (s *CloudStorageService) Create(req *CreateCloudStorageRequest) (*cloud_config_repo.CloudConfig, error) {
	config := &cloud_config_repo.CloudConfig{
		Uid:             req.Uid,
		Type:            req.Type,
		BucketName:      req.BucketName,
		AccountId:       req.AccountId,
		AccessKeyId:     req.AccessKeyId,
		AccessKeySecret: req.AccessKeySecret,
		CustomPath:      req.CustomPath,
	}

	id, err := s.cloudConfigDao.Create(config)
	if err != nil {
		return nil, err
	}

	return s.cloudConfigDao.GetByID(id)
}

// GetUserConfigs 获取用户的所有云存储配置
func (s *CloudStorageService) GetUserConfigs(uid int64) ([]*cloud_config_repo.CloudConfig, error) {
	return s.cloudConfigDao.GetByUID(uid)
}

// Update 更新云存储配置
func (s *CloudStorageService) Update(uid int64, id int64, req *CreateCloudStorageRequest) (*cloud_config_repo.CloudConfig, error) {
	config, err := s.cloudConfigDao.GetByID(id)
	if err != nil {
		return nil, err
	}

	if config.Uid != uid {
		return nil, errors.New("unauthorized")
	}

	config.Type = req.Type
	config.BucketName = req.BucketName
	config.AccountId = req.AccountId
	config.AccessKeyId = req.AccessKeyId
	config.AccessKeySecret = req.AccessKeySecret
	config.CustomPath = req.CustomPath

	if err := s.cloudConfigDao.Update(config); err != nil {
		return nil, err
	}

	return config, nil
}

// Delete 删除云存储配置
func (s *CloudStorageService) Delete(uid int64, id int64) error {
	return s.cloudConfigDao.Delete(uid, id)
}

// GetByID 根据ID获取配置
func (s *CloudStorageService) GetByID(id int64) (*cloud_config_repo.CloudConfig, error) {
	return s.cloudConfigDao.GetByID(id)
}

// GetUserConfigsByType 获取用户指定类型的所有配置
func (s *CloudStorageService) GetUserConfigsByType(uid int64, typeStr string) ([]*cloud_config_repo.CloudConfig, error) {
	return s.cloudConfigDao.GetByType(uid, typeStr)
}

// GetUserConfigByType 获取用户指定类型的单个配置
func (s *CloudStorageService) GetUserConfigByType(uid int64, typeStr string) (*cloud_config_repo.CloudConfig, error) {
	return s.cloudConfigDao.GetByUIDAndType(uid, typeStr)
}
