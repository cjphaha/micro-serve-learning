package ServicesImpl

import (
	"context"
	"github.com/go-playground/validator"
	"micro-toolbox/Database"
	"micro-toolbox/Services"
	"micro-toolbox/logs"
	"time"
)

type UserService struct {

}

var(
	validate = validator.New()
)

func (this *UserService) UserReg(ctx context.Context, req *Services.UserModel, rsp *Services.RegResponse) error {
	users := Database.User{UserName: req.UserName,UserId: int(req.UserId),UserPwd: req.UserPwd}
	logs.Log.Println("用户发送过来的数据为: ",users)
	users.UserDate = time.Now()
	err := validate.Struct(&users)
	if err != nil{
		errs := err.(validator.ValidationErrors)
		for _,v := range errs{
			logs.Log.Println("UserReg request参数错误: \n",v)
		}
	}
	db := Database.GetDb()
	err = db.Create(&users).Error
	if err != nil{
		logs.Log.Println(err.Error())
	}
	rsp.Message = ""
	rsp.Status = "success"
	return nil
}
