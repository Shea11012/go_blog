package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting,error) {
	vp := viper.New()
	vp.SetConfigName("config")
	basePath,_ := os.Getwd()
	for filepath.Base(basePath) != "go_blog" {
		basePath = filepath.Dir(basePath)
	}
	basePath = filepath.Join(basePath,"configs")
	vp.AddConfigPath(basePath)

	for _,config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}

	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil,err
	}
	s := &Setting{vp}
	s.WatchSettingChange()
	return s,nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}
