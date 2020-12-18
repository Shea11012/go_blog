package models

import (
	"errors"
	"gorm.io/gorm"
)

type Auth struct {
	*Model
	AppKey string `gorm:"type:varchar(20);default:''" json:"app_key"`
	AppSecret string `gorm:"type:varchar(50);default:''" json:"app_secret"`
}

func (a Auth) Get(db *gorm.DB) (*Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ?",a.AppKey,a.AppSecret)
	err := db.First(&auth).Error
	if err != nil && errors.Is(err,gorm.ErrRecordNotFound) {
		return nil,err
	}

	return &auth,nil
}