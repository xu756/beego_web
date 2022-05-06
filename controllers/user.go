package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"myweb/models"
	"strconv"
)

type Login struct {
	beego.Controller
}

func (request *Login) Post() {
	// 获取请求json参数
	user := make(map[string]string)
	data := request.Ctx.Input.RequestBody
	json.Unmarshal(data, &user)
	//返回设置
	res := make(map[string]string)

	o := orm.NewOrm()
	var userInfo models.User
	userInfo.Name = user["username"]
	userInfo.Password = user["password"]
	err := o.Read(&userInfo, "Name")
	if err != nil {
		res["status"] = strconv.Itoa(300)
		res["message"] = "用户不存在"
		request.Data["json"] = res
		request.ServeJSON()
		return
	}
	//使用md5加密
	h := md5.New()
	h.Write([]byte(user["password"]))
	password := hex.EncodeToString(h.Sum(nil))
	//判断密码是否正确
	fmt.Println(password)
	if userInfo.Password != password {
		res["status"] = strconv.Itoa(400)
		res["message"] = "密码错误"
		request.Data["json"] = res
		request.ServeJSON()
		return
	}
	res["status"] = strconv.Itoa(200)
	res["message"] = "登录成功"
	request.Data["json"] = res
	request.ServeJSON()
	return
}
