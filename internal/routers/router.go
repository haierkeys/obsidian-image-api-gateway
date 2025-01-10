package routers

import (
	"embed"
	"io/fs"
	"net/http"
	"time"

	_ "github.com/haierkeys/obsidian-image-api-gateway/docs"
	"github.com/haierkeys/obsidian-image-api-gateway/global"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/middleware"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/routers/apiRouter"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/limiter"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.BucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter(frontendFiles embed.FS) *gin.Engine {

	frontendAssets, _ := fs.Sub(frontendFiles, "frontend/assets")
	frontendIndexContent, _ := frontendFiles.ReadFile("frontend/index.html")

	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", frontendIndexContent)
	})
	r.StaticFS("/assets", http.FS(frontendAssets))
	r.NoRoute(middleware.NoFound())
	api := r.Group("/api")
	{
		api.Use(middleware.AppInfo())
		api.Use(gin.Logger())
		api.Use(middleware.RateLimiter(methodLimiters))
		api.Use(middleware.ContextTimeout(time.Duration(global.Config.App.DefaultContextTimeout) * time.Second))
		api.Use(middleware.Cors())
		api.Use(middleware.Translations())
		api.Use(middleware.AccessLog())
		if global.Config.Server.RunMode == "debug" {
			api.Use(gin.Recovery())
		} else {
			api.Use(middleware.Recovery())
		}
		// 对404 的处理
		// r.NoRoute(middleware.NoFound())
		// r.Use(middleware.Tracing())
		api.GET("/debug/vars", apiRouter.Expvar)
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		userApiR := api.Group("/user")
		{
			userApiR.POST("/register", apiRouter.NewUser().Register)
			userApiR.POST("/login", apiRouter.NewUser().Login)
			userApiR.Use(middleware.UserAuthToken()).POST("/cloud_config", apiRouter.NewCloudConfig().UpdateAndCreate)
			userApiR.Use(middleware.UserAuthToken()).DELETE("/cloud_config", apiRouter.NewCloudConfig().Delete)
			userApiR.Use(middleware.UserAuthToken()).GET("/cloud_config", apiRouter.NewCloudConfig().List)
			userApiR.Use(middleware.UserAuthToken()).POST("/upload", apiRouter.NewUpload().UserUpload)
		}
		api.Use(middleware.AuthToken()).POST("/upload", apiRouter.NewUpload().Upload)
	}
	if global.Config.LocalFS.IsEnabled && global.Config.LocalFS.HttpfsIsEnable {
		r.StaticFS(global.Config.LocalFS.SavePath, http.Dir(global.Config.LocalFS.SavePath))
	}

	return r
}
