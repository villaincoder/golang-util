package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
)

type BaseModel struct {
	RowId     uint64     `gorm:"primary_key"`
	Id        *string    `gorm:"not null;type:uuid;unique_index"`
	CreatedAt *time.Time `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"not null"`
	DeletedAt *time.Time
}

func NewBaseModel() BaseModel {
	nowTime := util.Time2Ptr(time.Now())
	return BaseModel{
		Id:        util.Str2Ptr(util.NewUUIDStr()),
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
	}
}

func MigrateModels(db *gorm.DB, dropExists bool, models ...interface{}) (err error) {
	for _, model := range models {
		if dropExists {
			err = db.DropTableIfExists(model).Error
			if err != nil {
				return
			}
		}
		err = db.AutoMigrate(model).Error
		if err != nil {
			return
		}
	}
	return
}

func CreateModels(db *gorm.DB, models ...interface{}) (err error) {
	for _, model := range models {
		err = db.Create(model).Error
		if err != nil {
			return
		}
	}
	return
}
