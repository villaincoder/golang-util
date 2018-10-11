package dbutil

import (
	"time"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
)

type BaseModel struct {
	RowId     uint64    `gorm:"primary_key"`
	Id        string    `gorm:"not null;type:uuid;unique_index"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt *time.Time
}

func NewBaseModel() BaseModel {
	nowTime := time.Now()
	return BaseModel{
		Id:        util.NewUUIDStr(),
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
	}
}

func (m BaseModel) ContainsId() bool {
	return m.RowId > 0 || m.Id != ""
}

func (m BaseModel) BeforeSave() (err error) {
	if m.ContainsId() {
		if m.Id == "" {
			err = errors.New("dbutil.BaseModel.Id empty")
		} else if util.IsInvalidTime(m.CreatedAt) {
			err = errors.New("dbutil.BaseModel.CreatedAt invalid")
		} else if util.IsInvalidTime(m.UpdatedAt) {
			err = errors.New("dbutil.BaseModel.UpdatedAt invalid")
		}
	}
	return
}

func MigrateModels(db *gorm.DB, models ...interface{}) (err error) {
	for _, model := range models {
		err = db.AutoMigrate(model).Error
		if err != nil {
			err = errors.WithStack(err)
			return
		}
	}
	return
}

func CreateModels(db *gorm.DB, models ...interface{}) (err error) {
	db = db.Begin()
	err = errors.WithStack(db.Error)
	if err != nil {
		return
	}
	for _, model := range models {
		err = db.Create(model).Error
		if err != nil {
			db.Rollback()
			err = errors.WithStack(err)
			return
		}
	}
	err = errors.WithStack(db.Commit().Error)
	if err != nil {
		db.Rollback()
	}
	return
}
