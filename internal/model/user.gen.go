// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/haierkeys/obsidian-image-api-gateway/pkg/timex"

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	UID       int64      `gorm:"column:uid;primaryKey" json:"uid" form:"uid"`
	Email     string     `gorm:"column:email;index:idx_pre_user_email,priority:1" json:"email" form:"email"`
	Username  string     `gorm:"column:username" json:"username" form:"username"`
	Password  string     `gorm:"column:password" json:"password" form:"password"`
	Salt      string     `gorm:"column:salt" json:"salt" form:"salt"`
	Token     string     `gorm:"column:token" json:"token" form:"token"`
	Avatar    string     `gorm:"column:avatar" json:"avatar" form:"avatar"`
	IsDeleted int64      `gorm:"column:is_deleted" json:"isDeleted" form:"isDeleted"`
	UpdatedAt timex.Time `gorm:"column:updated_at;type:datetime;autoUpdateTime" json:"updatedAt" form:"updatedAt"`
	CreatedAt timex.Time `gorm:"column:created_at;type:datetime;autoCreateTime" json:"createdAt" form:"createdAt"`
	DeletedAt timex.Time `gorm:"column:deleted_at;type:datetime;default:NULL" json:"deletedAt" form:"deletedAt"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
