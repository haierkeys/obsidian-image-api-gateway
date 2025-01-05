package app

import (
	"encoding/json"
	"time"

	"github.com/haierkeys/obsidian-image-api-gateway/global"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

type UserEntity struct {
	Uid      int64  `json:"uid"`
	Nickname string `json:"nickname"`
	time     int64  `json:"time"`
	IP       string `json:"ip"`
}

func ParseToken(token string) (*UserEntity, error) {

	var userEntity UserEntity

	resultStr, err := util.AuthDzCodeEncrypt(token, "DECODE", global.Config.Security.AuthTokenKey, 0)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resultStr), &userEntity)
	if err != nil {
		return nil, err
	}

	return &userEntity, nil

}

func GenerateToken(uid int64, nickname string, ip string, expiry int64) (string, error) {

	userEntity := UserEntity{
		Uid:      uid,
		Nickname: nickname,
		time:     time.Now().Unix(),
		IP:       ip,
	}

	userJson, err := json.Marshal(userEntity)
	if err != nil {
		return "", err
	}

	token, err := util.AuthDzCodeEncrypt(string(userJson), "ENCODE", global.Config.Security.AuthTokenKey, expiry)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUid(ctx *gin.Context) (out int64) {
	user, exist := ctx.Get("user_token")

	if exist == true {
		out = user.(*UserEntity).Uid
	}
	return
}

func GetIP(ctx *gin.Context) (out string) {
	user, exist := ctx.Get("user_token")
	if exist == true {
		out = user.(*UserEntity).IP
	}
	return out
}
