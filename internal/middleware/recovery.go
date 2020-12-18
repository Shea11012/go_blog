package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shea11012/go_blog/global"
	"github.com/shea11012/go_blog/pkg/app"
	"github.com/shea11012/go_blog/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				global.Logger.WithCallersFrames().Infof(c,s,err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()

		c.Next()
	}
}
