package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
	"web-gin/common/conf"
)

var DB *gorm.DB

func GetDB() {
	var err error
	DB, err = gorm.Open(conf.Config.Database.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Config.Database.User,
		conf.Config.Database.Password,
		conf.Config.Database.Host,
		conf.Config.Database.Name))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	DB.LogMode(true)             //开启sql输出日志
	DB.SingularTable(true)       //全局设置表名不可以为复数形式
	DB.DB().SetMaxIdleConns(10)  //用于设置闲置的连接数
	DB.DB().SetMaxOpenConns(100) //用于设置最大打开的连接数

	//DB.Callback().Create().Replace("gorm:update_time_stamp", CreatedAtCallback)
	//DB.Callback().Update().Replace("gorm:update_time_stamp", UpdatedAtCallback)
}

func CreatedAtCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func UpdatedAtCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}
}
