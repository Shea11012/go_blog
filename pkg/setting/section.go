package setting

import (
	"time"
)

var sections = make(map[string]interface{})

type ServerSetting struct {
	RunMode string
	HttpPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type AppSetting struct {
	DefaultPageSize int
	MaxPageSize int
	LogSavePath string
	LogFileName string
	LogFileExt string
	UploadSavePath string
	UploadServerUrl string
	UploadImageMaxSize int
	UploadImageAllowExts []string
}

type DatabaseSetting struct {
	DBType string
	Username string
	Password string
	Host string
	DBName string
	TablePrefix string
	Charset string
	ParseTime bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k,v)
	if err != nil {
		return err
	}

	if _,ok := sections[k];!ok {
		sections[k] = v
	}

	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k,v := range sections {
		err := s.ReadSection(k,v)
		if err != nil {
			return err
		}
	}

	return nil
}
