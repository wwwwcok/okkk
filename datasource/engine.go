package datasource

import (
	"fmt"
	"gg/day5_project/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func NewMysqlEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@/test1?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	engine.Sync2(new(model.Admin))
	fmt.Println("rrrrrrrrrrrrrrrrrrrrrrrrr")
	engine.ShowSQL(true)
	return engine
}
