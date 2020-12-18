package models

type ArticleTag struct {
	*Model
	TagID uint32 `gorm:"type:bigint unsigned" json:"tag_id"`
	ArticleID uint32 `gorm:"type:bigint unsigned" json:"article_id"`
}
