package dao

import (
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model/main_gen/user_repo"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/convert"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/timef"
)

type User struct {
	Uid       int64      `gorm:"column:uid;AUTO_INCREMENT" json:"uid" form:"uid"`                       //
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

// GetUserByUID 该结构体表示用户信息
// 包含用户的唯一标识符、头像、电子邮件、令牌、删除状态以及创建和更新时间
func (d *Dao) GetUserByUID(uid int64) (*User, error) {

	m, err := user_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &User{}).(*User), nil

}

// GetUserByEmail 根据电子邮件获取用户信息
func (d *Dao) GetUserByEmail(email string) (*User, error) {

	m, err := user_repo.NewQueryBuilder().
		WhereEmail(model.Eq, email).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &User{}).(*User), nil

}

// CreateMember 创建用户
func (d *Dao) CreateMember(dao *User) (int64, error) { // 修改参数类型为 User

	m := convert.StructAssign(dao, &user_repo.User{}).(*user_repo.User)

	id, err := m.Create()

	if err != nil {
		return 0, err
	}
	return id, nil
}

// CreateUser 创建用户
func (d *Dao) CreateUser(dao *User) (int64, error) { // 修改函数名为 CreateUser

	m := convert.StructAssign(dao, &user_repo.User{}).(*user_repo.User)

	id, err := m.Create()

	if err != nil {
		return 0, err
	}
	return id, nil
}
