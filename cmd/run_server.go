package cmd

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/haierkeys/obsidian-image-api-gateway/global"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/model"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/routers"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/logger"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/safe_close"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validatorV10 "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
)

type Server struct {
	logger            *zap.Logger // non-nil logger.
	httpServer        *http.Server
	privateHttpServer *http.Server
	sc                *safe_close.SafeClose
}

func NewServer(runEnv *runFlags) (*Server, error) {

	if err := global.ConfigLoad(runEnv.config); err != nil {
		return nil, err
	}

	runMode := runEnv.runMode
	if len(runMode) <= 0 {
		runMode = global.Config.Server.RunMode
	}

	if len(runMode) > 0 {
		gin.SetMode(runEnv.runMode)
	} else {
		gin.SetMode("release")
	}

	s := &Server{
		sc: safe_close.NewSafeClose(),
	}

	// Init logger.
	initLogger(s)

	initDatabase()

	initValidator()

	validator.RegisterCustom()

	// Start http api server
	if httpAddr := global.Config.Server.HttpPort; len(httpAddr) > 0 {
		s.logger.Info("starting api http server", zap.String("Config.Server.HttpPort", httpAddr))
		s.httpServer = &http.Server{
			Addr:           global.Config.Server.HttpPort,
			Handler:        routers.NewRouter(),
			ReadTimeout:    time.Duration(global.Config.Server.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(global.Config.Server.WriteTimeout) * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		s.sc.Attach(func(done func(), closeSignal <-chan struct{}) {
			defer done()
			errChan := make(chan error, 1)
			go func() {
				errChan <- s.httpServer.ListenAndServe()
			}()
			select {
			case err := <-errChan:
				s.logger.Error("private api http server err", zap.Error(err))
				s.sc.SendCloseSignal(err)
			case <-closeSignal:

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				// 停止HTTP服务器
				if err := s.httpServer.Shutdown(ctx); err != nil {
					s.logger.Error("service shutdown error", zap.Error(err))
				}

				// _ = s.httpServer.Close()
			}
		})
	}

	if httpAddr := global.Config.Server.PrivateHttpListen; len(httpAddr) > 0 {

		s.logger.Info("starting private api http server", zap.String("Config.Server.PrivateHttpListen", httpAddr))
		s.privateHttpServer = &http.Server{
			Addr:           global.Config.Server.PrivateHttpListen,
			Handler:        routers.NewPrivateRouter(),
			ReadTimeout:    time.Duration(global.Config.Server.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(global.Config.Server.WriteTimeout) * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		s.sc.Attach(func(done func(), closeSignal <-chan struct{}) {
			defer done()
			errChan := make(chan error, 1)
			go func() {
				errChan <- s.privateHttpServer.ListenAndServe()
			}()
			select {
			case err := <-errChan:
				s.logger.Error("private api http server err", zap.Error(err))
				s.sc.SendCloseSignal(err)
			case <-closeSignal:

				// _ = s.httpServer.Close()

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				// 停止HTTP服务器
				if err := s.privateHttpServer.Shutdown(ctx); err != nil {
					s.logger.Error("service shutdown error", zap.Error(err))
				}
			}
		})
	}

	return s, nil
}

func initLogger(s *Server) error {

	lg, err := logger.NewLogger(logger.Config{Level: global.Config.Log.Level, File: global.Config.Log.File, Production: global.Config.Log.Production})
	if err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}
	global.Logger = lg
	s.logger = lg

	return nil
}

func initValidator() error {
	global.Validator = validator.NewCustomValidator()
	global.Validator.Engine()
	binding.Validator = global.Validator

	var uni *ut.UniversalTranslator

	validate, ok := binding.Validator.Engine().(*validatorV10.Validate)
	if ok {

		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		uni = ut.New(en.New(), en.New(), zh.New())

		zhTran, _ := uni.GetTranslator("zh")
		enTran, _ := uni.GetTranslator("en")

		err := zh_translations.RegisterDefaultTranslations(validate, zhTran)
		if err != nil {
			return err
		}
		err = en_translations.RegisterDefaultTranslations(validate, enTran)
		if err != nil {
			return err
		}
	}

	global.Ut = uni

	return nil
}

func initDatabase() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.Config.Database)
	if err != nil {
		return err
	}

	// 自动迁移
	//db.AutoMigrate(&user_repo.User{})

	// 创建用户
	// user := cloud_config_repo.CloudConfig{}
	// user.CreatedAt = timef.Now()
	// db.Model(&user).Create(&user)
	// dump.P(user)
	// // 查询用户

	// db := global.DBEngine
	// m, err := cloud_config_repo.NewQueryBuilder().
	// 	WhereId(model.Eq, 1).
	// 	WhereIsDeleted(model.Eq, 0).
	// 	First()

	// m.CreatedAt = timef.Now()
	// //m.UpdatedAt = timef.Now()

	// db.Model(m).Updates(m)

	// dump.P(m)
	// dump.P(m.CreatedAt.Value())

	// 输出用户创建时间
	//fmt.Println("用户创建时间:", queriedUser.CreatedAt)

	return nil
}
