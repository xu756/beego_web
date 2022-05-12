package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// User 用户表
type User struct {
	Id       int       //`orm:"pk;auto"` 主键，自增
	Name     string    //用户名
	Password string    `orm:"size(255)"`                   //密码
	Email    string    `orm:"size(50)"`                    //邮箱
	Avatar   string    `orm:"size(255)"`                   //头像
	Created  time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
	Updated  time.Time `orm:"auto_now;type(datetime)"`     //更新时间
	Token    string    `orm:"size(255)"`                   //token
	//
}

func init() {
	orm.RegisterDataBase("default", "mysql", "xu:xjx756756@tcp(121.5.132.57:5700)/Beego_db?charset=utf8&loc=Local", 30)
	orm.RegisterModel(new(User)) //注册表
	orm.Debug = true
	orm.RunSyncdb("default", false, false) //同步数据库
}
