package cloud_config_repo

import "github.com/haierkeys/obsidian-image-api-gateway/pkg/timef"

//
//go:generate gormgen -structs CloudConfig -input . -pre pre_
type CloudConfig struct {
	Id              int64      `gorm:"column:id;primary_key;auto_increment" json:"id" form:"id"`                          //
	Uid             int64      `gorm:"column:uid;index;default:0" json:"uid" form:"uid"`                                  //
	Type            string     `gorm:"column:type;default:''" json:"type" form:"type"`                                    //
	BucketName      string     `gorm:"column:bucket_name;default:''" json:"bucketName" form:"bucketName"`                 //
	AccountId       string     `gorm:"column:account_id;default:''" json:"accountId" form:"accountId"`                    //
	AccessKeyId     string     `gorm:"column:access_key_id;default:''" json:"accessKeyId" form:"accessKeyId"`             //
	AccessKeySecret string     `gorm:"column:access_key_secret;default:''" json:"accessKeySecret" form:"accessKeySecret"` //
	CustomPath      string     `gorm:"column:custom_path;default:''" json:"customPath" form:"customPath"`                 //
	IsDeleted       int64      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                     //
	UpdatedAt       timef.Time `gorm:"column:updated_at;time;default:NULL" json:"updatedAt" form:"updatedAt"`             //
	CreatedAt       timef.Time `gorm:"column:created_at;time;default:NULL" json:"createdAt" form:"createdAt"`             //
	DeletedAt       timef.Time `gorm:"column:deleted_at;time;default:NULL" json:"deletedAt" form:"deletedAt"`             //
}
