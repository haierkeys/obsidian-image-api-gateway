package apiRouter

import (
	"github.com/haierkeys/obsidian-image-api-gateway/global"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/service"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/app"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/code"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CloudConfig struct {
}

func NewCloudConfig() *CloudConfig {
	return &CloudConfig{}
}

func (t *CloudConfig) UpdateAndCreate(c *gin.Context) {
	params := &service.CloudConfigRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, params)
	if !valid {
		global.Logger.Error("apiRouter.CloudConfig.UpdateAndCreate.BindAndValid errs: %v", zap.Error(errs))
		response.ToResponse(code.ErrorInvalidParams.WithDetails(errs.ErrorsToString()).WithData(errs.MapsToString()))
		return
	}
	uid := app.GetUid(c)
	if uid == 0 {
		global.Logger.Error("apiRouter.CloudConfig.UpdateAndCreate err uid=0")
		response.ToResponse(code.ErrorNotUserAuthToken)
		return
	}
	svc := service.New(c)
	err := svc.CloudConfigUpdateAndCreate(uid, params)
	if err != nil {
		global.Logger.Error("apiRouter.CloudConfig.UpdateAndCreate svc UpdateAndCreate err: %v", zap.Error(err))
		response.ToResponse(code.Failed.WithDetails(err.Error()))
		return
	}
	if params.Id == 0 {
		response.ToResponse(code.SuccessCreate)
	} else {
		response.ToResponse(code.SuccessUpdate)
	}
}

func (t *CloudConfig) Delete(c *gin.Context) {
	param := service.DeleteCloudConfigRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error("apiRouter.CloudConfig.Delete.BindAndValid svc Delete err: %v", zap.Error(errs))
		response.ToResponse(code.ErrorInvalidParams.WithDetails(errs.ErrorsToString()).WithData(errs.MapsToString()))
		return
	}
	uid := app.GetUid(c)
	if uid == 0 {
		global.Logger.Error("apiRouter.CloudConfig.Delete err uid=0")
		response.ToResponse(code.ErrorNotUserAuthToken)
		return
	}
	svc := service.New(c)
	err := svc.CloudConfigDelete(uid, &param)
	if err != nil {
		global.Logger.Error("apiRouter.CloudConfig.Delete svc Delete err: %v", zap.Error(err))
		response.ToResponse(code.Failed.WithDetails(err.Error()))
		return
	}
	response.ToResponse(code.SuccessDelete)
}

func (t *CloudConfig) List(c *gin.Context) {
	response := app.NewResponse(c)
	uid := app.GetUid(c)
	if uid == 0 {
		global.Logger.Error("apiRouter.CloudConfig.List err uid=0")
		response.ToResponse(code.ErrorNotUserAuthToken)
		return
	}
	svc := service.New(c)
	list, total, err := svc.CloudConfigList(uid, &app.Pager{Page: 1, PageSize: 10})
	if err != nil {
		response.ToResponse(code.Failed.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(code.Success, list, total)
}
