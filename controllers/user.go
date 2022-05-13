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
		res["code"] = strconv.Itoa(300)
		res["message"] = "用户不存在"
		request.Data["json"] = res
		request.ServeJSON()
		return
	}
	h := sha256.New()
	h.Write([]byte(user["password"]))
	password := hex.EncodeToString(h.Sum(nil))
	if userInfo.Password != password {
		res["code"] = strconv.Itoa(400)
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
	//todo: 设置缓存
	//beego.GlobalStorage.Set(token, userInfo, time.Hour*24)
	// 返回结果
	res["code"] = strconv.Itoa(200)
	res["message"] = "登录成功"
	request.Data["json"] = res
	request.ServeJSON()
	return
}

type Userinfo struct {
	beego.Controller
	token string
}

// Post 获取用户信息
func (request *Userinfo) Post() {
	// 获取token
	token := request.Ctx.Input.Param(":token")
	if token == "" {
		request.Data["json"] = map[string]string{"code": "300", "message": "token不能为空"}
		request.ServeJSON()
	}
	// 获取用户信息
	o := orm.NewOrm()
	var userInfo models.User
	userInfo.Token = token
	err := o.Read(&userInfo, "Token")
	if err != nil {
		request.Data["json"] = map[string]string{"code": "300", "message": "token不存在"}
		request.ServeJSON()
	}
	request.Data["json"] = map[string]interface{}{"code": "200", "message": "获取成功", "userinfo": userInfo}
	request.ServeJSON()
}

// Get 判断用户是否本人
func (request *Userinfo) Get() {
	// 获取token
	token := request.Ctx.Input.Param(":token")
	if token == "" {
		request.Data["json"] = map[string]string{"code": "300", "message": "请重新登录"}
		request.ServeJSON()
	}
	o := orm.NewOrm()
	var userInfo models.User
	userInfo.Token = token
	err := o.Read(&userInfo, "Token")
	if err != nil {
		request.Data["json"] = map[string]string{"code": "300", "message": "请重新登录"}
		request.ServeJSON()
	}
	request.Data["json"] = map[string]interface{}{"code": "200", "message": "是本人登录"}
	request.ServeJSON()
}
