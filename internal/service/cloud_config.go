package service

import (
	"github.com/haierkeys/obsidian-image-api-gateway/internal/dao"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model/main_gen/cloud_config_repo"
)

type CloudConfigService struct {
	dao *dao.CloudConfigDao
}

func NewCloudConfigService() *CloudConfigService {
	return &CloudConfigService{
		dao: dao.NewCloudConfigDao(),
	}
}

// 创建云存储配置
func (s *CloudConfigService) Create(param *dao.CloudConfig) (int64, error) {
	return s.dao.Create(param)
}

// 更新云存储配置
func (s *CloudConfigService) Update(param *dao.CloudConfig) error {
	return s.dao.Update(param)
}

// 获取用户的云存储配置列表
func (s *CloudConfigService) GetListByUid(uid int64) ([]*cloud_config_repo.CloudConfig, error) {
	return s.dao.GetListByUid(uid)
}

// 获取配置详情
func (s *CloudConfigService) GetById(id int64) (*cloud_config_repo.CloudConfig, error) {
	return s.dao.GetById(id)
}

// 删除配置
func (s *CloudConfigService) Delete(id int64) error {
	return s.dao.Delete(id)
}

// 检查配置是否属于用户
func (s *CloudConfigService) CheckOwnership(id int64, uid int64) (bool, error) {
	config, err := s.GetById(id)
	if err != nil {
		return false, err
	}
	if config == nil {
		return false, nil
	}
	return config.Uid == uid, nil
}
