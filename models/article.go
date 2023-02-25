package models

import (
	//"time"

	//"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	Title 	   string `json:"title"`
	Journal    string `json:"journal"`
	Author       string `json:"author"`
	Authors  string `json:"authors"`
	Date string `json:"date"`
	Link      string    `json:"link"`
	Papercode      string    `json:"papercode"`
	Abstract      string    `json:"abstract"`
	Theyear      string 	`json:"theyear"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}
func ExistArticleByTitle(id int ,title string) bool {
	var article Article
	db.Select("id").Where("title = ?", title).First(&article)

	if article.ID > 0 && article.ID != id{
		return true
	}

	return false
}
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	

	return
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		Title:     data["title"].(string),
		Journal:      data["journal"].(string),
		Author: data["author"].(string),
		Date: data["date"].(string),
		Authors: data["authors"].(string),
		Papercode: data["papercode"].(string),
		Link: data["link"].(string),
		Abstract: data["abstract"].(string),
		Theyear:data["theyear"].(string),

		
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}


