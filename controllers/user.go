package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"myweb/models"
)

type Login struct {
	beego.Controller
}

func (request *Login) Post() {
	user := make(map[string]string)
	data := request.Ctx.Input.RequestBody
	json.Unmarshal(data, &user)
	//获取json数据
	fmt.Println(user["username"])
	request.Data["json"] = user
	request.ServeJSON()
	//查询数据库
	createuser(user["username"], user["password"])
	//o := orm.NewOrm()
	//var userInfo models.User
	//userInfo.Name = user["username"]
	//err := o.Read(&userInfo, "Name")
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	fmt.Println("查询成功")
	//	fmt.Println(userInfo.Password)
	//}
}

//创建用户
func createuser(username string, password string) bool {
	o := orm.NewOrm()
	var userInfo models.User
	userInfo.Name = username
	userInfo.Password = password
	o.Insert(&userInfo)
	return true
}
