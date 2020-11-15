package Database

import (
	"github.com/go-playground/validator"
	"micro-toolbox/logs"
	"testing"
	"time"
)

func TestValidate(t *testing.T) {
	type User struct {
		//gorm.Model
		UserId int
		UserName string `validate:"required,min=3,max=18"`
		UserPwd string  `validate:"required,min=10,max=18"`
		UserDate time.Time `gorm:"column:user_date"`
	}
	v := validator.New()
	user := User{}
	user.UserName = "h"
	var err  error
	if err = v.Struct(&user) ; err != nil {
		errs := err.(validator.ValidationErrors)
		for _,v := range errs{
			logs.Log.Println("TestValidate: \n",v)
		}

	}
	t.Log(err)
}
