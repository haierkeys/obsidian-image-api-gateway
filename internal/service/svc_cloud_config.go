package service

import (
	"github.com/haierkeys/obsidian-image-api-gateway/internal/dao"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/app"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/code"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/convert"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/storage"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/timex"
)

type CloudConfig struct {
	ID              int64      `json:"id"`              // ID
	Type            string     `json:"type"`            // 类型
	BucketName      string     `json:"bucketName"`      // 存储桶名称
	Endpoint        string     `json:"endpoint"`        // 端点
	Region          string     `json:"region"`          // 区域
	AccountID       string     `json:"accountId"`       // 账户ID
	AccessKeyID     string     `json:"accessKeyId"`     // 访问密钥ID
	AccessKeySecret string     `json:"accessKeySecret"` // 访问密钥秘密
	CustomPath      string     `json:"customPath"`      // 自定义路径
	AccessURLPrefix string     `json:"accessUrlPrefix"` // 访问地址前缀
	IsEnabled       int64      `json:"isEnabled"`       // 是否启用
	UpdatedAt       timex.Time `json:"updatedAt"`       // 更新时间
	CreatedAt       timex.Time `json:"createdAt"`       // 创建时间
}

type CloudConfigRequest struct {
	ID              int64  `form:"id"`                                                // ID
	Type            string `form:"type" binding:"required,gte=1"`                     // 类型
	Endpoint        string `form:"endpoint"`                                          // 端点 oss
	Region          string `form:"region"`                                            // 区域 s3
	AccountID       string `form:"accountId"`                                         // 账户ID r2
	BucketName      string `form:"bucketName"`                                        // 存储桶名称
	AccessKeyID     string `form:"accessKeyId"`                                       // 访问密钥ID
	AccessKeySecret string `form:"accessKeySecret"`                                   // 访问密钥秘密
	CustomPath      string `form:"customPath"`                                        // 自定义路径
	AccessURLPrefix string `form:"accessUrlPrefix"  binding:"required,min=2,max=100"` // 访问地址前缀
	IsEnabled       int64  `form:"isEnabled"`                                         // 是否启用
}

type DeleteCloudConfigRequest struct {
	Id int64 `form:"id" binding:"required,gte=1"`
}

// CloudTypeList 方法用于获取云存储类型列表
func (svc *Service) CloudTypeEnabledList() ([]storage.CloudType, error) {
	return storage.GetIsUserEnabledStorageTypes(), nil
}

// CloudConfigList 方法用于获取指定用户的云存储配置列表
func (svc *Service) CloudConfigList(uid int64, pager *app.Pager) ([]*CloudConfig, int, error) {

	// 统计指定用户的云存储配置数量
	count, err := svc.dao.CountListByUID(uid)
	if err != nil {
		return nil, 0, err // 如果发生错误，返回 nil 和错误信息
	}

	// 获取指定用户的云存储配置列表
	daoList, err := svc.dao.GetListByUID(pager.Page, pager.PageSize, uid)
	if err != nil {
		return nil, 0, err // 如果发生错误，返回 nil 和错误信息
	}

	var list []*CloudConfig
	// 将获取到的配置转换为 CloudConfig 类型并添加到列表中
	for _, m := range daoList {
		list = append(list, convert.StructAssign(m, &CloudConfig{}).(*CloudConfig))
	}

	// 返回配置列表和数量
	return list, int(count), nil
}

// 云存储管理 - 更新云存储配置的方法
func (svc *Service) CloudConfigUpdateAndCreate(uid int64, params *CloudConfigRequest) (int64, error) {

	// 检查云存储类型是否有效
	if !storage.StorageTypeMap[params.Type] {
		return 0, code.ErrorInvalidStorageType
	}

	// 检查云存储类型是否启用
	if err := storage.IsUserEnabled(params.Type); err != nil {
		return 0, err
	}

	//云存储内容设置项检查
	if storage.CloudStorageTypeMap[params.Type] {
		if params.BucketName == "" {
			return 0, code.ErrorInvalidCloudStorageBucketName
		}
		if params.AccessKeyID == "" {
			return 0, code.ErrorInvalidCloudStorageAccessKeyID
		}
		if params.AccessKeySecret == "" {
			return 0, code.ErrorInvalidCloudStorageAccessKeySecret
		}
	}

	// 检查云存储类型是否为 r2
	if params.Type == storage.R2 {

		// 检查账户ID是否为空
		if params.AccountID == "" {
			return 0, code.ErrorInvalidCloudStorageAccountId
		}

	} else if params.Type == storage.S3 {
		// 检查区域是否为空
		if params.Region == "" {
			return 0, code.ErrorInvalidCloudStorageRegion
		}
	} else if params.Type == storage.OSS {
		// 检查端点是否为空
		if params.Endpoint == "" {
			return 0, code.ErrorInvalidCloudStorageEndpoint
		}
	} else if params.Type == storage.MinIO {
		// 检查端点是否为空
		if params.Endpoint == "" {
			return 0, code.ErrorInvalidCloudStorageEndpoint
		}
	}

	// 调用数据访问层的更新方法
	da := convert.StructAssign(params, &dao.CloudConfigSet{}).(*dao.CloudConfigSet)

	var id int64
	var err error
	if params.ID == 0 {
		id, err = svc.dao.Create(da, uid)
		if err != nil {
			// 如果发生错误，返回错误信息
			return 0, err
		}
		svc.dao.DisableBatch(id, uid)
	} else {
		id = params.ID
		err := svc.dao.Update(da, params.ID, uid)
		if err != nil {
			// 如果发生错误，返回错误信息
			return 0, err
		}
		if params.IsEnabled == 1 {
			svc.dao.DisableBatch(params.ID, uid)
		}
	}
	return id, nil
}

// 删除指定用户的云存储配置
func (svc *Service) CloudConfigDelete(uid int64, param *DeleteCloudConfigRequest) error {
	// 调用数据访问层的删除方法
	err := svc.dao.Delete(param.Id, uid)
	if err != nil {
		// 如果发生错误，返回错误信息
		return err
	}
	// 返回 nil 表示删除成功
	return nil
}
