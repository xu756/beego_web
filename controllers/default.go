package controllers

import beego "github.com/beego/beego/v2/server/web"

// Index 处理函数
type Index struct {
	beego.Controller
}

func (request *Index) Get() {
	request.Ctx.WriteString("Hello World")
}
