package models

import (
	"github.com/shea11012/go_blog/pkg/app"
)

type Article struct {
	*Model
	Title string `gorm:"type:varchar(100);default:''" json:"title"`
	Desc string `gorm:"type:varchar(255);default:''" json:"desc"`
	CoverImageUrl string `gorm:"type:varchar(255);default''" json:"cover_image_url"`
	Context string `gorm:"type:longtext" json:"context"`
	State uint8 `gorm:"type:tinyint unsigned" json:"state"`
}

type ArticleSwagger struct {
	List []*Article
	Pager *app.Pager
}