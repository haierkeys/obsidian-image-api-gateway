package service

import (
	"github.com/haierkeys/obsidian-image-api-gateway/internal/dao"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/timef"
)

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

type CreateUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

// UserRegister 用户注册
func (svc *Service) UserRegister(param *CreateUserRequest) (int64, error) {
	user := &dao.User{
		Email: param.Email,
		// 其他字段可以根据需要设置，例如头像等
	}

	id, err := svc.dao.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UserLogin 用户登录
func (svc *Service) UserLogin(param *LoginUserRequest) (*dao.User, error) {
	user, err := svc.dao.GetUserByCredentials(param.Email, param.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
