package upload

import (
	"github.com/shea11012/go_blog/global"
	"github.com/shea11012/go_blog/pkg/setting"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	loadSetting()
	m.Run()
}

func loadSetting() {
	settings,err := setting.NewSetting()
	global.CheckError(err)

	err = settings.ReadSection("Server",&global.ServerSetting)
	global.CheckError(err)

	err = settings.ReadSection("App",&global.AppSetting)
	global.CheckError(err)

	err = settings.ReadSection("Database",&global.DatabaseSetting)
	global.CheckError(err)

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
}

func TestCheckContainExt(t *testing.T) {
	fileName := "f26ef914245883c80f181c4aade2ed04.jpg"
	if b := CheckContainExt(TypeImage,fileName);!b {
		t.Fatalf("文件名后缀检测失败")
	}
}
