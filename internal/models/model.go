package models

import (
	"fmt"
	"github.com/shea11012/go_blog/global"
	"github.com/shea11012/go_blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Model struct {
	ID 	uint32	`gorm:"primaryKey;type:bigint unsigned" json:"id"`
	CreatedBy string `gorm:"type:varchar(100)" json:"created_by"`
	ModifiedBy string `gorm:"type:varchar(100)" json:"modified_by"`
	CreatedAt *time.Time `gorm:"type:datetime" json:"created_at"`
	ModifiedAt *time.Time `gorm:"type:datetime" json:"modified_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime" json:"deleted_at"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	db,err := gorm.Open(mysql.Open(s),&gorm.Config{})
	global.CheckError(err)

	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}

	sqlDb,err := db.DB()
	global.CheckError(err)

	db.Callback().Create().Before("gorm:create").Register("update_created_at",updateTimeForCreateCallback)

	sqlDb.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDb.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db,nil
}

func updateTimeForCreateCallback(db *gorm.DB) {
	field := db.Statement.Schema.LookUpField("CreatedAt")
	if field != nil {
		if _,isZero := field.ValueOf(db.Statement.ReflectValue);isZero {
			// err := field.Set(db.Statement.ReflectValue,time.Now())
			// if err != nil {
			// 	global.Logger.Fatalf("updateTimeForCreateCallback err：%v",err)
			// }
		}
	}
}


