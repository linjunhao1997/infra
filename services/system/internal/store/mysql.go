package store

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func initGormDB() {
	dsn := `root:123456@tcp(192.168.174.129:3306)/test?charset=utf8&parseTime=True&loc=Local`

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	GormDB = db
}

func init() {
	initGormDB()
}
