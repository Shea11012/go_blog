package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shea11012/go_blog/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRow int64 `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{ctx}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	r.Ctx.JSON(http.StatusOK,data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int64) {
	r.Ctx.JSON(http.StatusOK,gin.H{
		"list":list,
		"pager":Pager{
			Page: GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRow: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code":err.Code,
		"msg": err.Msg,
	}
	details := err.Details

	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(http.StatusOK,response)
}


