package models

import (
	"hello-world/libs"

	"github.com/jinzhu/gorm"
)

type ChaptersModel struct {
	gorm.Model

	//ID       uint    `gorm:"primary_key;auto_increment"`
	Id_season uint   `gorm:"type:int(10);not null;"`
	Image     string `gorm:"type:varchar(300);not null;"`
	Resourse  string `gorm:"type:varchar(300);not null;"`
	Name      string `gorm:"type:varchar(300);not null;"`
	Order     uint   `gorm:"type:int(10);not null;"`
}

func (ChaptersModel) TableName() string {
	return "chapters"
}
func (this *ChaptersModel) Objvoid() ChaptersModel {
	return ChaptersModel{}
}
func (t *ChaptersModel) ListAllByIdSeason(id uint) []ChaptersModel {
	fields := new(ChaptersModel)

	fields.Id_season = id
	return t.ListAllFor(fields)
}
func (t *ChaptersModel) ChapterById(id uint) ChaptersModel {
	fields := new(ChaptersModel)
	fields.ID = id
	return t.GetBy(fields)
}

func (this *ChaptersModel) ListAllFor(fields *ChaptersModel) []ChaptersModel {
	var data = []ChaptersModel{}
	var err error
	if fields == nil {
		err = libs.DB.Find(&data).Error
	} else {
		err = libs.DB.Where(fields).Order("id ASC").Find(&data).Error
	}
	if err != nil {
		return []ChaptersModel{}
	}
	return data
}
func (this *ChaptersModel) GetBy(fields *ChaptersModel) ChaptersModel {
	var data = this.Objvoid()
	err := libs.DB.Where(fields).Find(&data).Error
	if err != nil {
		return this.Objvoid()
	}
	return data
}
func (this *ChaptersModel) Add(Id_season uint, Name string, imagen string, resource string, order uint) (ChaptersModel, error) {
	var data ChaptersModel
	data.Id_season = Id_season
	data.Image = imagen
	data.Name = Name
	data.Resourse = resource
	data.Order = order
	if err := libs.DB.Create(&data).Error; err != nil {
		return this.Objvoid(), err
	} else {
		return data, nil
	}

}
func (this *ChaptersModel) Del(id uint) error {
	var data ChaptersModel
	if err := libs.DB.Where("id = ?", id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
