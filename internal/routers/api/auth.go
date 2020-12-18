package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shea11012/go_blog/global"
	"github.com/shea11012/go_blog/internal/service"
	"github.com/shea11012/go_blog/pkg/app"
	"github.com/shea11012/go_blog/pkg/errcode"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if !valid {
		global.Logger.Infof(c,"app.bindandvalid errs :%v",errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Infof(c,"svc checkAuth err: %v",err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token,err := app.GenerateToken(param.AppKey,param.AppSecret)
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token":token,
	})
}
