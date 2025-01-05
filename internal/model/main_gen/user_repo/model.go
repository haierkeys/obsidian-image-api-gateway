package user_repo

import "github.com/haierkeys/obsidian-image-api-gateway/pkg/timef"

//
//go:generate gormgen -structs User -input . -pre pre_
type User struct {
	Uid       int64      `gorm:"column:uid;primary_key;auto_increment" json:"uid" form:"uid"`           //
	Email     string     `gorm:"column:email;default:''" json:"email" form:"email"`                     //
	Username  string     `gorm:"column:username;default:''" json:"username" form:"username"`            //
	Password  string     `gorm:"column:password;default:''" json:"password" form:"password"`            //
	Salt      string     `gorm:"column:salt;default:''" json:"salt" form:"salt"`                        //
	Token     string     `gorm:"column:token;default:''" json:"token" form:"token"`                     //
	Avatar    string     `gorm:"column:avatar;default:''" json:"avatar" form:"avatar"`                  //
	IsDeleted int64      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`         //
	UpdatedAt timef.Time `gorm:"column:updated_at;time;default:NULL" json:"updatedAt" form:"updatedAt"` //
	CreatedAt timef.Time `gorm:"column:created_at;time;default:NULL" json:"createdAt" form:"createdAt"` //
	DeletedAt timef.Time `gorm:"column:deleted_at;time;default:NULL" json:"deletedAt" form:"deletedAt"` //
}
