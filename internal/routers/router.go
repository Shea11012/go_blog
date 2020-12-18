package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/shea11012/go_blog/docs"
	"github.com/shea11012/go_blog/global"
	"github.com/shea11012/go_blog/internal/middleware"
	"github.com/shea11012/go_blog/internal/routers/api"
	v1 "github.com/shea11012/go_blog/internal/routers/api/v1"
	"github.com/shea11012/go_blog/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(middleware.Tracing())
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		r.Use(middleware.AccessLog(),middleware.Recovery())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use(middleware.Translations())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		tag := v1.NewTag()
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.GET("/tags", tag.List)

		article := v1.NewArticle()
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
