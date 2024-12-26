package user_repo

import (
	"github.com/haierspi/golang-image-upload-service/pkg/timef"
)

// User 会员表
type User struct {
	Uid       int64      `gorm:"column:uid;primary_key;AUTO_INCREMENT" json:"uid"`                 // 用户UID
	Avatar    string     `gorm:"column:avatar;default:''" json:"avatar"`                           // 头像
	Email     string     `gorm:"column:email;unique;default:''" json:"email"`                      // 邮箱
	Token     string     `gorm:"column:token;default:''" json:"token"`                             // 用户TOKEN
	IsDeleted int32      `gorm:"column:is_deleted;default:0" json:"isDeleted"`                     // 是否删除
	UpdatedAt timef.Time `gorm:"column:updated_at;default:'0000-00-00 00:00:00'" json:"updatedAt"` // 更新时间
	CreatedAt timef.Time `gorm:"column:created_at;default:'0000-00-00 00:00:00'" json:"createdAt"` // 创建时间
	DeletedAt timef.Time `gorm:"column:deleted_at;default:'0000-00-00 00:00:00'" json:"deletedAt"` // 标记删除时间
}
