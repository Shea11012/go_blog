package dao

import (
	"github.com/shea11012/go_blog/internal/models"
)

func (d *Dao) GetAuth(appKey, appSecret string) (*models.Auth, error) {
	auth := models.Auth{AppKey: appKey,AppSecret: appSecret}
	return auth.Get(d.engine)
}
