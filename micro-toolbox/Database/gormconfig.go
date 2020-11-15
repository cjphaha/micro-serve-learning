package Database

import (
	"github.com/Unknwon/goconfig"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"micro-toolbox/logs"
)

var (
	cfg, _ = goconfig.LoadConfigFile("./config.ini");
	name, _      = cfg.GetValue("mysql", "username");
	password,_ = cfg.GetValue("mysql", "password");
	url,_ = cfg.GetValue("mysql", "url");
	DB *gorm.DB
)

func init(){
	db, err := gorm.Open("mysql",name+":"+password+"@tcp"+url+"?charset=utf8&parseTime=True&loc=Local")
	//db,err := gorm.Open("mysql","root:123@tcp(127.0.0.1:3307)/dbname?charset=utf8")
	if err != nil{
		logs.Log.Println(err)
	}
	//SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.DB().SetMaxOpenConns(50)   //设置数据库连接池最大连接数
	db.DB().SetMaxIdleConns(10)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	DB = db
	AutoMerage()
}

func GetDb() *gorm.DB{
	return DB
}
