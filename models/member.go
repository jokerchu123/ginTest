package models

import (
	//"time"

	//"github.com/jinzhu/gorm"
)
type Member struct {
	Model

	Identity       string `json:"identity"`
	Name     string `json:"name"`
	Phone      string 	`json:"phone"`
	Achievement       string `json:"achievement"`
	Mail     string `json:"mail"`
	Research      string 	`json:"research"`
	Introduction      string 	`json:"introduction"`
}

func GetMembers(pageNum int, pageSize int, maps interface{}) (member []Member) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&member)

	return
}
func GetMember(id int) (member Member) {
	db.Where("id = ?", id).First(&member)
	

	return
}
func GetMemberTotal(maps interface{}) (count int) {
	db.Model(&Member{}).Where(maps).Count(&count)
	return
}
func ExistMemberByName(id int ,name string) bool {
	var member Member
	db.Select("id").Where("name = ?", name).First(&member)
	if member.ID > 0 && member.ID != id{
		return true
	}

	return false
}

func AddMember(name string,phone string,achievement string,mail string,research string,identity string,introduction string) bool {
	db.Create(&Member{
		Name:		name,
		Identity:   identity,
		Phone:		phone,
		Mail:		mail,
		Achievement:   achievement,
		Research:		research,
		Introduction: introduction,
	})

	return true
}
func DeleteMember(id int) bool {
	db.Where("id = ?", id).Delete(&Member{})

	return true
}

func ExistMemberByID(id int) bool {
	var member Member
	db.Select("id").Where("id = ?", id).First(&member)
	if member.ID > 0 {
		return true
	}

	return false
}


func EditMember(id int, data interface{}) bool {
	db.Model(&Member{}).Where("id = ?", id).Updates(data)

	return true
}
