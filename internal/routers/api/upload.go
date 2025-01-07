package api

import (
	"github.com/haierkeys/obsidian-image-api-gateway/global"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/service"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/app"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/code"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/fsutil"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// Upload 上传文件
func (u Upload) Upload(c *gin.Context) {

	params := &service.ClientUploadParams{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, params)

	if !valid {
		global.Logger.Error("app.BindAndValid errs: %v", zap.Error(errs))
		response.ToResponse(code.ErrorInvalidParams.WithDetails(errs.ErrorsToString()).WithData(errs.MapsToString()))
		return
	}

	var svcUploadFileData *service.FileInfo
	var svc = service.New(c)
	var err error

	file, fileHeader, errf := c.Request.FormFile("imagefile")
	defer file.Close()

	if errf != nil {
		global.Logger.Error("app.ErrorInvalidParams len 0", zap.Error(errf))
		response.ToResponse(code.ErrorInvalidParams)
	}

	svcUploadFileData, err = svc.UploadFile(fsutil.TypeImage, file, fileHeader, params)
	if err != nil {
		global.Logger.Error("svc.UploadFile err: %v", zap.Error(err))
		response.ToResponse(code.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(code.Success.WithData(svcUploadFileData))

	return

}
