package global

import (
	"github.com/shea11012/go_blog/pkg/logger"
	"github.com/shea11012/go_blog/pkg/setting"
)

var (
	ServerSetting *setting.ServerSetting
	AppSetting *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	Logger *logger.Logger
	JWTSetting *setting.JWTSetting
)
