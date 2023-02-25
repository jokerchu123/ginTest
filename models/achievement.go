package models

import (
	//"time"

	//"github.com/jinzhu/gorm"
)
type Achievement struct {
	Model

	Name       string `json:"name"`
	Address     string `json:"address"`
	Category	string `json:"category"`
}

func GetAchievements(pageNum int, pageSize int, maps interface{}) (achievement []Achievement) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&achievement)

	return
}
func GetAchievement(id int) (achievement Achievement) {
	db.Where("id = ?", id).First(&achievement)
	

	return
}
func GetAchievementTotal(maps interface{}) (count int) {
	db.Model(&Achievement{}).Where(maps).Count(&count)

	return
}
func ExistAchievementByName(id int ,name string) bool {
	var achievement Achievement
	db.Select("id").Where("name = ?", name).First(&achievement)
	if achievement.ID > 0 && achievement.ID != id{
		return true
	}

	return false
}

func AddAchievement(name string,address string,category string) bool {
	db.Create(&Achievement{
		Name:		name,
		Address:    address,
		Category: 	category,
	})

	return true
}
func DeleteAchievement(id int) bool {
	db.Where("id = ?", id).Delete(&Achievement{})

	return true
}

func ExistAchievementByID(id int) bool {
	var achievement Achievement
	db.Select("id").Where("id = ?", id).First(&achievement)
	if achievement.ID > 0 {
		return true
	}

	return false
}


func EditAchievement(id int, data interface{}) bool {
	db.Model(&Achievement{}).Where("id = ?", id).Updates(data)

	return true
}
