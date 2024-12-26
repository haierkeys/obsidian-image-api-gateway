package user_repo

import "github.com/haierspi/golang-image-upload-service/pkg/timef"

//
//go:generate gormgen -structs User -input . -pre pre_
type User struct {
	Uid       int64      `gorm:"column:uid;AUTO_INCREMENT" json:"uid" form:"uid"`                       //
	Avatar    string     `gorm:"column:avatar;default:''" json:"avatar" form:"avatar"`                  //
	Email     string     `gorm:"column:email;default:''" json:"email" form:"email"`                     //
	Token     string     `gorm:"column:token;default:''" json:"token" form:"token"`                     //
	IsDeleted int64      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`         //
	UpdatedAt timef.Time `gorm:"column:updated_at;time;default:NULL" json:"updatedAt" form:"updatedAt"` //
	CreatedAt timef.Time `gorm:"column:created_at;time;default:NULL" json:"createdAt" form:"createdAt"` //
	DeletedAt timef.Time `gorm:"column:deleted_at;time;default:NULL" json:"deletedAt" form:"deletedAt"` //
}
