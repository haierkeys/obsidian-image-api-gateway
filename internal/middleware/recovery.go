package middleware

import (
	"fmt"
	"time"

	"github.com/haierkeys/obsidian-image-api-gateway/global"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/app"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/code"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/email"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	var defailtMailer *email.Email
	if global.Config.Email.ErrorReportEnable {
		defailtMailer = email.NewEmail(&email.SMTPInfo{
			Host:     global.Config.Email.Host,
			Port:     global.Config.Email.Port,
			IsSSL:    global.Config.Email.IsSSL,
			UserName: global.Config.Email.UserName,
			Password: global.Config.Email.Password,
			From:     global.Config.Email.From,
		})
	}

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				global.Log().Fatal("panic recover err", zap.String("err", fmt.Sprintf("%v", err)))

				if global.Config.Email.ErrorReportEnable {
					err := defailtMailer.SendMail(
						global.Config.Email.To,
						fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
						fmt.Sprintf("错误信息: %v", err),
					)
					if err != nil {
						global.Log().Panic("mail.SendMail err", zap.Error(err))
					}
				}

				app.NewResponse(c).ToResponse(code.ErrorServerInternal)
				c.Abort()
			}
		}()
		c.Next()
	}
}
