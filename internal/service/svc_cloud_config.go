package service

import (
	"github.com/haierkeys/obsidian-image-api-gateway/internal/dao"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/app"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/convert"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/timex"
)

type CloudConfig struct {
	Id              int64      `json:"Id"`
	Type            string     `json:"type"`
	BucketName      string     `json:"bucketName"`
	AccountId       string     `json:"accountId"`
	AccessKeyId     string     `json:"accessKeyId"`
	AccessKeySecret string     `json:"accessKeySecret"`
	CustomPath      string     `json:"customPath"`
	AccessUrlPrefix string     `json:"accessUrlPrefix"`
	IsEnabled       int64      `json:"isEnabled"`
	UpdatedAt       timex.Time `json:"updatedAt"`
	CreatedAt       timex.Time `json:"createdAt"`
}

type CloudConfigRequest struct {
	Id              int64  `form:"Id"`                                                // ID
	Type            string `form:"type" binding:"required,gte=1"`                     // 类型
	BucketName      string `form:"bucketName" binding:"required,gte=1"`               // 存储桶名称
	AccountId       string `form:"accountId"`                                         // 账户ID
	AccessKeyId     string `form:"accessKeyId" binding:"required,min=2,max=100"`      // 访问密钥ID
	AccessKeySecret string `form:"accessKeySecret" binding:"required,min=2,max=100"`  // 访问密钥秘密
	CustomPath      string `form:"customPath"`                                        // 自定义路径
	AccessUrlPrefix string `form:"accessUrlPrefix"  binding:"required,min=2,max=100"` // 访问地址前缀
	IsEnabled       int64  `form:"isEnabled"`                                         // 是否启用
}

type DeleteCloudConfigRequest struct {
	Id int64 `form:"Id" binding:"required,gte=1"`
}

// CloudConfigList 方法用于获取指定用户的云存储配置列表
func (svc *Service) CloudConfigList(uid int64, pager *app.Pager) ([]*CloudConfig, int, error) {

	// 统计指定用户的云存储配置数量
	count, err := svc.dao.CountListByUid(uid)
	if err != nil {
		return nil, 0, err // 如果发生错误，返回 nil 和错误信息
	}

	// 获取指定用户的云存储配置列表
	daoList, err := svc.dao.GetListByUid(pager.Page, pager.PageSize, uid)
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
func (svc *Service) CloudConfigUpdateAndCreate(uid int64, params *CloudConfigRequest) error {

	// if params.Type == OSS {
	// 	return code.ErrorInvalidParams
	// }
	// 调用数据访问层的更新方法
	da := convert.StructAssign(params, &dao.CloudConfigSet{}).(*dao.CloudConfigSet)
	if params.Id == 0 {
		id, err := svc.dao.Create(da, uid)
		if err != nil {
			// 如果发生错误，返回错误信息
			return err
		}
		svc.dao.DisableBatch(id, uid)
	} else {
		err := svc.dao.Update(da, params.Id, uid)
		if err != nil {
			// 如果发生错误，返回错误信息
			return err
		}
		if params.IsEnabled == 1 {
			svc.dao.DisableBatch(params.Id, uid)
		}
	}
	return nil
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
