package models

import (
	"github.com/shea11012/go_blog/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name string `gorm:"type:varchar(32)" json:"name"`
	State uint8 `gorm:"type:unsigned tinyint" json:"state"`
}

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?",t.Name)
	}

	db = db.Where("state = ?",t.State)
	err := db.Model(&t).Count(&count).Error
	if err != nil {
		return 0,err
	}

	return count,nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag,error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if t.Name != "" {
		db = db.Where("name = ?",t.Name)
	}

	db = db.Where("state = ?",t.State)
	if err = db.Find(&tags).Error;err != nil {
		return nil, err
	}

	return tags,nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB,data map[string]interface{}) error {
	return db.Model(&t).Where("id = ?",t.ID).Updates(data).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Delete(&t).Error
}