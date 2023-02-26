package main

import (
	"context"
	"fmt"
	"infra/model"
	"infra/pkg/dao"
	"infra/services/system/pkg/useraccount"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/* type user struct {
	Id    int    `gorm:"column:id"`
	Test1 string `gorm:"column:test1"`
	Test2 string `gorm:"column:test2"`
}

func (u *user) TableName() string {
	return "test"
} */

func main() {

	dsn := `root:123456@tcp(192.168.174.129:3306)/test?charset=utf8&parseTime=True&loc=Local`

	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	db.AutoMigrate(useraccount.Admin{}, useraccount.Sysadmin{}, useraccount.Role{})

	/* err = db.Create(&model.User{Name: "dog2", Groups: model.UserGroups{{Name: "toy1"}, {Name: "toy2"}}}).Error
	if err != nil {
		panic(err)
	} */

	usrDao := dao.NewDAO[model.User](db)

	/* total, err := usrDao.Count(context.Background(), dao.Filter{
		Eq: []dao.Eq{{ColumnName: "name", ColumnValue: "dog5"}},
	})

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(total) */

	/* usrDao.Create(context.Background(), &model.User{
		Name: "dog10",
	}) */

	/* if _, err := usrDao.Upsert(context.Background(), []string{"pwd", "mtime"}, []string{"name"}, &model.User{Pwd: "123456", Name: "dog5"}); err != nil {
		panic(err)
	} */

	/* res, err := usrDao.FindPage(context.Background(), dao.Query{Pagination: dao.Pagination{CurrentPage: 1, PageSize: 1}, Filter: dao.Filter{
		Eq: []dao.Eq{{ColumnName: "disabled", ColumnValue: 0}},
	}}, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", res) */

	/* res, err := usrDao.Find(context.Background(), dao.Filter{Eq: []dao.Eq{{ColumnName: "disabled", ColumnValue: 1}}},
		&dao.Sorter{Columns: []dao.OrderBy{{ColumnName: "id", Desc: false}}})
	if err != nil {
		panic(err)
	}

	for _, e := range res {
		fmt.Println(e.Id)
	} */

	/* res, err := usrDao.FindById(context.Background(), 13)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res) */

	/* res, err := usrDao.FindByIds(context.Background(), 1, 13)
	if err != nil {
		panic(err)
	}
	for _, e := range res {
		fmt.Printf("%+v", e)

	} */

	count, err := usrDao.DeleteByIds(context.Background(), 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	/* res, err := usrDao.UpdateByColumnMap(context.Background(), 7, map[string]interface{}{"email": "852857605@qq.com"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res) */

	/* res, err := usrDao.UpdateByModel(context.Background(), 7, &model.User{Email: "12345678@qq.com"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res) */
}
