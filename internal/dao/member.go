package dao

import (
	"github.com/haierspi/golang-image-upload-service/internal/model"
	"github.com/haierspi/golang-image-upload-service/internal/model/db_gen/user_repo"
	"github.com/haierspi/golang-image-upload-service/pkg/convert"
	"github.com/haierspi/golang-image-upload-service/pkg/timef"
)

type Member struct {
	Uid           int64      `gorm:"column:uid;primary_key;AUTO_INCREMENT" json:"uid" form:"uid"`                            // 用户UID
	Nickname      string     `gorm:"column:nickname;default:''" json:"nickname" form:"nickname"`                             // 用户昵称
	Avatar        string     `gorm:"column:avatar;default:''" json:"avatar" form:"avatar"`                                   // 头像
	Mobile        string     `gorm:"column:mobile;default:''" json:"mobile" form:"mobile"`                                   // 电话号码
	Name          string     `gorm:"column:name;default:''" json:"name" form:"name"`                                         // 真实姓名
	Idcard        string     `gorm:"column:idcard;default:''" json:"idcard" form:"idcard"`                                   // 身份证号
	IsValidate    int32      `gorm:"column:is_validate;default:0" json:"isValidate" form:"isValidate"`                       // 是否身份验证
	Openid        string     `gorm:"column:openid;default:''" json:"openid" form:"openid"`                                   // openid
	Unionid       string     `gorm:"column:unionid;default:''" json:"unionid" form:"unionid"`                                // unionid
	Gender        string     `gorm:"column:gender;default:''" json:"gender" form:"gender"`                                   // 性别
	Language      string     `gorm:"column:language;default:''" json:"language" form:"language"`                             // 语言
	City          string     `gorm:"column:city;default:''" json:"city" form:"city"`                                         // 城市
	Province      string     `gorm:"column:province;default:''" json:"province" form:"province"`                             // 省份
	Country       string     `gorm:"column:country;default:''" json:"country" form:"country"`                                // 国家
	AvatarUrl     string     `gorm:"column:avatar_url;default:''" json:"avatarUrl" form:"avatarUrl"`                         // 微信的头像
	SessionKey    string     `gorm:"column:session_key;default:''" json:"sessionKey" form:"sessionKey"`                      // session_key
	Token         string     `gorm:"column:token;default:''" json:"token" form:"token"`                                      // 用户TOKEN
	WeixinToken   string     `gorm:"column:weixin_token;default:''" json:"weixinToken" form:"weixinToken"`                   // 用户微信TOKEN
	ChangeNameNum int32      `gorm:"column:change_name_num;default:0" json:"changeNameNum" form:"changeNameNum"`             // 用户昵称剩余修改次数
	HannelsId     int64      `gorm:"column:hannels_id;default:0" json:"hannelsId" form:"hannelsId"`                          // 渠道ID
	IsDeleted     int32      `gorm:"column:is_deleted;default:0" json:"isDeleted" form:"isDeleted"`                          // 是否删除
	UpdatedAt     timef.Time `gorm:"column:updated_at;time;default:'0000-00-00 00:00:00'" json:"updatedAt" form:"updatedAt"` // 更新时间
	CreatedAt     timef.Time `gorm:"column:created_at;time;default:'0000-00-00 00:00:00'" json:"createdAt" form:"createdAt"` // 创建时间
	DeletedAt     timef.Time `gorm:"column:deleted_at;time;default:'0000-00-00 00:00:00'" json:"deletedAt" form:"deletedAt"` // 标记删除时间
	AppId         int64      `gorm:"column:app_id;default:0" json:"appId" form:"appId"`                                      // 注册来源 应用的ID
}

func (d *Dao) GetMemberByUID(uid int64) (*Member, error) {

	m, err := user_repo.NewQueryBuilder().
		WhereUid(model.Eq, uid).
		WhereIsDeleted(model.Eq, 0).
		First()

	if err != nil {
		return nil, err
	}

	return convert.StructAssign(m, &Member{}).(*Member), nil

}

func (d *Dao) GetMemberByMobile(mobile string) (*Member, error) {

	m, err := user_repo.NewQueryBuilder().
		WhereMobile(model.Eq, mobile).
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
