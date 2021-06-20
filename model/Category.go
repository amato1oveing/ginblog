package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在
func CheckCate(name string) (code int) {
	var cates Category
	db.Select("id").Where("name = ?", name).First(&cates)
	if cates.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

//新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//todo查看分类下的所有文章

//查询分类列表
func GetCate(pageSize int, pageNum int) []Category {
	var cates []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cates
}

//编辑分类
func EditCate(id int, data *Category) int {
	var cates Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cates).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除分类
func DeleteCate(id int) int {
	var cates Category
	err := db.Where("id = ?", id).Delete(&cates).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
