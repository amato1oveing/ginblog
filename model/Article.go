package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

//新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//todo查看文章下的所有文章

//todo查询单个文章

//查询文章列表
func GetArt(pageSize int, pageNum int) []Article {
	var art []Article
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return art
}

//编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&art).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除文章
func DeleteArt(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
