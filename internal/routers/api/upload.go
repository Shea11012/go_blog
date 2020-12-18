package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shea11012/go_blog/internal/service"
	"github.com/shea11012/go_blog/pkg/app"
	"github.com/shea11012/go_blog/pkg/convert"
	"github.com/shea11012/go_blog/pkg/errcode"
	"github.com/shea11012/go_blog/pkg/upload"
)

type Upload struct {

}

func NewUpload() Upload {
	return Upload{}
}

// @Summary 上传文件
// @Param file formData string true "上传文件"
// @Param type formData int true "上传文件类型"
func (u Upload) UploadFile(c *gin.Context)  {
	response := app.NewResponse(c)
	file,fileHeader,err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errRsp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	if fileHeader == nil || fileType < 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo,err := svc.UploadFile(upload.FileType(fileType),file,fileHeader)
	if err != nil {
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
