package models

import (
	"hello-world/libs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SeasonModel struct {
	gorm.Model
	//ID       uint    `gorm:"primary_key;auto_increment"`
	Chapters uint   `gorm:"type:int(10);not null;"`
	Image    string `gorm:"type:varchar(300);not null;"`
	Name     string `gorm:"type:varchar(300);not null;"`
}

func (SeasonModel) TableName() string {
	return "seasons"
}
func (this *SeasonModel) Objvoid() SeasonModel {
	return SeasonModel{}
}
func (t *SeasonModel) ListAll() []SeasonModel {
	return t.ListAllFor(nil)
}
func (t *SeasonModel) SeasonById(id uint) SeasonModel {
	fields := new(SeasonModel)
	fields.ID = id
	return t.GetBy(fields)
}

func (this *SeasonModel) ListAllFor(fields *SeasonModel) []SeasonModel {
	var data = []SeasonModel{}
	var err error
	if fields == nil {
		err = libs.DB.Find(&data).Error
	} else {
		err = libs.DB.Where(fields).Order("id desc").Find(&data).Error
	}
	if err != nil {
		return []SeasonModel{}
	}
	return data
}

func (this *SeasonModel) GetBy(fields *SeasonModel) SeasonModel {
	var data = this.Objvoid()
	err := libs.DB.Where(fields).Find(&data).Error
	if err != nil {
		return this.Objvoid()
	}
	return data
}

func (this *SeasonModel) Get(id uint) SeasonModel {
	var data = this.Objvoid()
	err := libs.DB.Where("id = ? ", id).Find(&data).Error
	if err != nil {
		return this.Objvoid()
	}
	return data
}
func (this *SeasonModel) Add(Name string, imagen string) (SeasonModel, error) {
	var data SeasonModel
	data.Chapters = 0
	data.Image = imagen
	data.Name = Name
	if err := libs.DB.Create(&data).Error; err != nil {
		return this.Objvoid(), err
	} else {
		return data, nil
	}

}
func (this *SeasonModel) Del(id uint) error {
	var data SeasonModel
	if err := libs.DB.Where("id = ?", id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
