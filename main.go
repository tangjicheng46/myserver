package main

import (
	_ "myserver/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:abcd1234@tcp(localhost:3306)/myserver?charset=utf8")
}

func main() {
	beego.Run()
}
