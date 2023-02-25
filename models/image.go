package models

import (
	//"time"

	//"github.com/jinzhu/gorm"
)
type Image struct {
	Model
	Name       string `json:"name"`
	Date       string `json:"date"`
	Address     string `json:"address"`
}

func GetImages(pageNum int, pageSize int, maps interface{}) (image []Image) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&image)

	return
}
func GetImage(id int) (image Image) {
	db.Where("id = ?", id).First(&image)
	

	return
}
func GetImageTotal(maps interface{}) (count int) {
	db.Model(&Image{}).Where(maps).Count(&count)

	return
}
func ExistImageByName(id int ,name string) bool {
	var image Image
	db.Select("id").Where("name = ?", name).First(&image)
	if image.ID > 0 && image.ID != id{
		return true
	}

	return false
}

func AddImage(name string,date string,address string) bool {
	db.Create(&Image{
		Name:		name,
		Date:      date,
		Address:address,
	})

	return true
}
func DeleteImage(id int) bool {
	db.Where("id = ?", id).Delete(&Image{})

	return true
}

func ExistImageByID(id int) bool {
	var image Image
	db.Select("id").Where("id = ?", id).First(&image)
	if image.ID > 0 {
		return true
	}

	return false
}


func EditImage(id int, data interface{}) bool {
	db.Model(&Image{}).Where("id = ?", id).Updates(data)

	return true
}
