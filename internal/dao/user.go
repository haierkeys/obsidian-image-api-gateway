package dao

import (
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model/main_gen/user_repo"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/convert"
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

func (d *Dao) GetUserByEmail(email string) (*Member, error) {

	m, err := user_repo.NewQueryBuilder().
		WhereEmail(model.Eq, email).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &Member{}).(*Member), nil

}

func (d *Dao) CreateMember(dao *Member) (int64, error) {

	m := convert.StructAssign(dao, &user_repo.Member{}).(*user_repo.Member)

	id, err := m.Create()

	if err != nil {
		return 0, err
	}
	return id, nil
}
