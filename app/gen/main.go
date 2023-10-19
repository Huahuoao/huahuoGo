package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:123456@tcp(43.136.99.219:3309)/calculator?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      ".././app/dao/dal",
		ModelPkgPath: ".././app/dao/model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
