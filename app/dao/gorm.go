package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	dsn := "root:123456@tcp(43.136.99.219:3309)/calculator?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		log.Fatalln("db connection err")
	}
}

//func QueryData(k, j, c int) []YourModel {
//	db := // 这里是从数据库连接池中获取一个可用的 *gorm.DB 对象
//
//	queryWrapper := db.Model(&YourModel{})
//	if k == 1 {
//		queryWrapper = queryWrapper.Where("a = ?", k)
//	}
//	if j == 1 {
//		queryWrapper = queryWrapper.Where("b = ?", j)
//	}
//	if c == 1 {
//		queryWrapper = queryWrapper.Order("g DESC")
//	}
//	results := make([]YourModel, 0)
//	queryWrapper.Find(&results)
//
//	return results
//}
