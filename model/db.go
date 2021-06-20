package model

import (
	"fmt"
	"ginblog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Println("连接数据库失败，err:", err)
	}
	db.SingularTable(true) //禁用复数
	db.AutoMigrate(&User{}, &Category{}, &Article{})

	db.DB().SetMaxIdleConns(10)                  //设置数据池最大限制连接数
	db.DB().SetMaxOpenConns(100)                 //设置数据库最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 10) //设置连接最长复用时间

}
