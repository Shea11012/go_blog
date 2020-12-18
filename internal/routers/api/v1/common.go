package v1

import (
	"github.com/shea11012/go_blog/pkg/app"
	"github.com/shea11012/go_blog/pkg/errcode"
)

func checkValid(valid bool, errs app.ValidErrors) *errcode.Error {
	if !valid {
		return errcode.InvalidParams.WithDetails(errs.Errors()...)
	}
	return nil
}
