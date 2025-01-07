package service

import (
    "github.com/haierkeys/obsidian-image-api-gateway/internal/dao"
    "github.com/haierkeys/obsidian-image-api-gateway/pkg/app"
    "github.com/haierkeys/obsidian-image-api-gateway/pkg/code"
    "github.com/haierkeys/obsidian-image-api-gateway/pkg/convert"
    "github.com/haierkeys/obsidian-image-api-gateway/pkg/timex"
    "github.com/haierkeys/obsidian-image-api-gateway/pkg/util"
)

type User struct {
    Uid       int64      `gorm:"column:uid;AUTO_INCREMENT" json:"uid" form:"uid"`
    Username  string     `gorm:"column:username;default:''" json:"username" form:"username"`            //
    Avatar    string     `gorm:"column:avatar;default:''" json:"avatar" form:"avatar"`                  //
    Email     string     `gorm:"column:email;default:''" json:"email" form:"email"`                     //
    Token     string     `gorm:"column:token;default:''" json:"token" form:"token"`                     //
    UpdatedAt timex.Time `gorm:"column:updated_at;time;default:NULL" json:"updatedAt" form:"updatedAt"` //
    CreatedAt timex.Time `gorm:"column:created_at;time;default:NULL" json:"createdAt" form:"createdAt"` //
}

type UserCreateRequest struct {
    Email    string `json:"email" form:"email" binding:"required,email"`
    Password string `json:"password" form:"password" binding:"required"`
}

type UserLoginRequest struct {
    Email    string `form:"email" binding:"required,email"`
    Password string `form:"password" binding:"required"`
}

type UserRegisterSendEmail struct {
    Email string `json:"email" form:"email" binding:"required,email"`
}

func (svc *Service) UserRegisterSendEmail(param *UserCreateRequest) (int64, error) {
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

// UserRegister 用户注册
func (svc *Service) UserRegister(param *UserCreateRequest) (*User, error) {

    emailUser, err := svc.dao.GetUserByEmail(param.Email)
    if emailUser != nil {
        return nil, code.ErrorUserEmailAlreadyExists
    }

    password, err := util.GeneratePasswordHash(param.Password)
    if err != nil {
        return nil, code.ErrorPasswordNotValid
    }

    user := &dao.User{
        Email:    param.Email,
        Password: password,
        // 其他字段可以根据需要设置，例如头像等
    }

    id, err := svc.dao.CreateUser(user)
    if err != nil {
        return nil, err
    }

    expiry := 30 * 24 * 60 * 60
    ip := svc.ctx.ClientIP()
    userAuthToken, err := app.GenerateToken(id, "", ip, int64(expiry))
    user.Token = userAuthToken

    return convert.StructAssign(user, &User{}).(*User), nil
}

// UserLogin 用户登录
func (svc *Service) UserLogin(param *UserLoginRequest) (*User, error) {

    user, err := svc.dao.GetUserByEmail(param.Email)
    if err != nil {
        return nil, code.ErrorUserNotFound
    }

    if !util.CheckPasswordHash(user.Password, param.Password) {
        return nil, code.ErrorUserLoginPasswordFailed
    }

    expiry := 30 * 24 * 60 * 60
    ip := svc.ctx.ClientIP()
    userAuthToken, err := app.GenerateToken(user.Uid, user.Username, ip, int64(expiry))
    user.Token = userAuthToken

    return convert.StructAssign(user, &User{}).(*User), nil
}
