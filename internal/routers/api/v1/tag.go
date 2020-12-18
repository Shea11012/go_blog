package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shea11012/go_blog/global"
	"github.com/shea11012/go_blog/internal/service"
	"github.com/shea11012/go_blog/pkg/app"
	"github.com/shea11012/go_blog/pkg/convert"
	"github.com/shea11012/go_blog/pkg/errcode"
)

type Tag struct {}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context)  {

}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxLength(100)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} models.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags [get]
func (t Tag) List(c *gin.Context)  {
	param := service.TagListRequest{}

	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if errRsp:= checkValid(valid,errs); errRsp != nil {
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c),PageSize: app.GetPageSize(c)}
	totalRaw,err := svc.CountTag(&service.CountTagRequest{
		Name: param.Name,
		State: param.State,
	})

	if err != nil {
		global.Logger.Fatalf("svc.CountTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags,err := svc.GetTagList(&param,pager)
	if err != nil {
		global.Logger.Fatalf("svc.GetTagList err: %v",err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags,totalRaw)
	return
}

// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名称" minLength(3) maxLength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} models.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags [post]
func (t Tag) Create(c *gin.Context)  {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if errRsp := checkValid(valid,errs);errRsp != nil {
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	if err := svc.CreateTag(&param); err != nil {
		global.Logger.Fatalf("svc.CreateTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce json
// @Param id path int true "标签ID"
// @Param name body string true "标签名称" minLength(3) maxLength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} models.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags/{id} [put]
func (t Tag) Update(c *gin.Context)  {
	params := service.UpdateTagRequest{
		ID:convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&params)
	if errRsp := checkValid(valid,errs);errRsp != nil {
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		global.Logger.Infof(c,"svc.UpdateTag err: %v",errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	if err := svc.UpdateTag(&params); err != nil {
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce json
// @Param id formData int true "标签ID"
// @Success 200 {object} models.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags [delete]
func (t Tag) Delete(c *gin.Context)  {
	param := service.DeleteTagRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	if errRsp := checkValid(valid,errs); errRsp != nil {
		global.Logger.Fatalf("tag.DeleteTag err: %v",errRsp)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	svc := service.New(c.Request.Context())
	if err := svc.DeleteTag(&param); err != nil {
		global.Logger.Fatalf("svc.DeleteTag err: %v",err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}