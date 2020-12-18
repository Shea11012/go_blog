package dao

import (
	"github.com/shea11012/go_blog/internal/models"
	"github.com/shea11012/go_blog/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tag := models.Tag{Name: name,State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*models.Tag, error) {
	tag := models.Tag{
		Name: name,
		State: state,
	}
	pageOffset := app.GetPageOffset(page,pageSize)

	return tag.List(d.engine,pageOffset,pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := models.Tag{
		Model: &models.Model{CreatedBy: createdBy},
		Name:  name,
		State: state,
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := models.Tag{
		Name: name,
		State: state,
		Model:&models.Model{ModifiedBy: modifiedBy,ID: id},
	}
	data := map[string]interface{}{
		"name":name,
		"state":state,
		"modified_by": modifiedBy,
	}
	return tag.Update(d.engine,data)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := models.Tag{
		Model:&models.Model{ID: id},
	}

	return tag.Delete(d.engine)
}
