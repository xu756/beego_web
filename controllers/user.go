package controllers

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"myweb/models"
	"strconv"
	"time"
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
	h := sha256.New()
	h.Write([]byte(user["password"]))
	password := hex.EncodeToString(h.Sum(nil))
	if userInfo.Password != password {
		res["status"] = strconv.Itoa(400)
		res["message"] = "密码错误"
		request.Data["json"] = res
		request.ServeJSON()
		return
	}
	h = sha512.New()
	h.Write([]byte(user["password"] + user["username"] + time.Now().String()))
	token := hex.EncodeToString(h.Sum(nil))
	// 设置token
	res["token"] = token
	userInfo.Token = token
	o.Update(&userInfo, "Token")
	//使用缓存
	//beego.GlobalStorage.Set(token, userInfo, time.Hour*24)
	// 返回结果
	res["status"] = strconv.Itoa(200)
	res["message"] = "登录成功"
	request.Data["json"] = res
	request.ServeJSON()
	return
}
