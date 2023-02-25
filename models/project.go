package models

import (
	//"time"

	//"github.com/jinzhu/gorm"
)
type Project struct {
	Model

	Name       string `json:"name"`
	Link     string `json:"link"`
	Theyear      string 	`json:"theyear"`
}

func GetProjects(pageNum int, pageSize int, maps interface{}) (project []Project) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&project)

	return
}
func GetProject(id int) (project Project) {
	db.Where("id = ?", id).First(&project)
	

	return
}
func GetProjectTotal(maps interface{}) (count int) {
	db.Model(&Project{}).Where(maps).Count(&count)

	return
}
func ExistProjectByName(id int ,name string) bool {
	var project Project
	db.Select("id").Where("name = ?", name).First(&project)
	if project.ID > 0 && project.ID != id{
		return true
	}

	return false
}

func AddProject(name string,link string,theyear string) bool {
	db.Create(&Project{
		Name:		name,
		Link:      link,
		Theyear:theyear,
	})

	return true
}
func DeleteProject(id int) bool {
	db.Where("id = ?", id).Delete(&Project{})

	return true
}

func ExistProjectByID(id int) bool {
	var project Project
	db.Select("id").Where("id = ?", id).First(&project)
	if project.ID > 0 {
		return true
	}

	return false
}


func EditProject(id int, data interface{}) bool {
	db.Model(&Project{}).Where("id = ?", id).Updates(data)

	return true
}
