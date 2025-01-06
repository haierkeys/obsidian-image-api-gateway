package routers

import (
	"net/http"
	"time"

	_ "github.com/haierkeys/obsidian-image-api-gateway/docs"
	"github.com/haierkeys/obsidian-image-api-gateway/global"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/middleware"
	"github.com/haierkeys/obsidian-image-api-gateway/internal/routers/api"
	"github.com/haierkeys/obsidian-image-api-gateway/pkg/limiter"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
	// gin-swagger middleware
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.BucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.AppInfo())
	r.Use(gin.Logger())

	if global.Config.Server.RunMode == "debug" {
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.Recovery())
	}
	// 对404 的处理
	r.NoRoute(middleware.NoFound())
	// r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(time.Duration(global.Config.App.DefaultContextTimeout) * time.Second))
	r.Use(middleware.Cors())
	r.Use(middleware.Translations())

	// r.Use(middleware.AuthToken())

	r.GET("/debug/vars", api.Expvar)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiR := r.Group("/api")
	if global.Config.Server.RunMode != "debug" {
		apiR.Use(middleware.AccessLog())
	}

	// a.Use(middleware.AuthToken())
	// a.POST("/upload", api.NewUpload().Upload)
	userApiR := apiR.Group("/user")
	{
		userApiR.POST("/register", api.NewUser().Register)
		userApiR.POST("/login", api.NewUser().Login)
		userApiR.Use(middleware.UserAuthToken()).POST("/cloud_config", api.NewCloudConfig().UpdateAndCreate)
		userApiR.Use(middleware.UserAuthToken()).DELETE("/cloud_config", api.NewCloudConfig().Delete)
		userApiR.Use(middleware.UserAuthToken()).GET("/cloud_config", api.NewCloudConfig().List)
	}

	apiR.Use(middleware.AuthToken()).POST("/upload", api.NewUpload().Upload)

	// .Use(middleware.UserAuthToken())

	if global.Config.LocalFS.Enable && global.Config.LocalFS.HttpfsEnable {
		r.StaticFS(global.Config.LocalFS.SavePath, http.Dir(global.Config.LocalFS.SavePath))
	}

	return r
}
