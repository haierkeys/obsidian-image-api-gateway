package app

import (
	"encoding/json"

	"github.com/haierkeys/obsidian-image-api-gateway/global"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

type UserEntity struct {
	Uid    int64  `json:"uid"`
	Expiry int64  `json:"expiry"`
	IP     string `json:"ip"`
}

func ParseToken(token string) (*UserEntity, error) {

	var userEntity UserEntity

	resultStr, err := util.AuthDzCodeEncrypt(token, "DECODE", global.Config.Security.AuthTokenKey, 0)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resultStr), &userEntity)

	if err == nil {
		return &userEntity, nil
	}

	return nil, err

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

func GetExpiration(ctx *gin.Context) (out int64) {
	user, exist := ctx.Get("user_token")
	if exist == true {
		out = user.(*UserEntity).Expiry
	}
	return out
}
