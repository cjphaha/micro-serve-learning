package Database

import (
	"time"
)

type User struct {
	//gorm.Model
	UserId int `gorm:"column:user_id;PRIMARY_KEY";`
	UserName string `gorm:"column:user_name;type:varchar(50)",validate:"required";`
	UserPwd string  `gorm:"column:user_pwd;type:varchar(50)" validate:"required,min=10";`
	UserDate time.Time `gorm:"column:user_date"`
}

func AutoMerage(){
	db := GetDb()
	//db.CreateTable(&User{})
	isHave := db.HasTable(&User{})
	if !isHave{
		db.CreateTable(&User{})
	}
	db.AutoMigrate(&User{})
}