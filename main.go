package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "myweb/models"  // 导入models包  	数据库
	_ "myweb/routers" // 导入routers包		路由
)

func main() {
	beego.Run()
}
