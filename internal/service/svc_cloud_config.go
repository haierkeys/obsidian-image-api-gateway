package service

import (
	"github.com/gookit/goutil/dump"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/dao"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/app"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/code"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/convert"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/storage"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/timex"
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
	UpdatedAt       timex.Time `json:"updatedAt" form:"updatedAt"`             // 更新时间，自动更新时间戳
	CreatedAt       timex.Time `json:"createdAt" form:"createdAt"`             // 创建时间，自动创建时间戳
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
	User            string `form:"user"`                                              // 访问用户名
	Password        string `form:"password"`                                          // 密码
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
	} else if params.Type == storage.WebDAV {
		// 检查端点是否为空
		if params.Endpoint == "" {
			return 0, code.ErrorWebDAVInvalidEndpoint
		}
		if params.User == "" {
			return 0, code.ErrorWebDAVInvalidUser
		}
		if params.Password == "" {
			return 0, code.ErrorWebDAVInvalidPassword
		}
	}
	if params.AccessURLPrefix == "" {
		return 0, code.ErrorInvalidAccessURLPrefix
	}

	// 调用数据访问层的更新方法
	da := convert.StructAssign(params, &dao.CloudConfigSet{}).(*dao.CloudConfigSet)

	dump.P(uid)
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
