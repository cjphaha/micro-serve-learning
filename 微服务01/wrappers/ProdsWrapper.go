package wrappers

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	"miniserver/Services"
	"strconv"
)

//降级方法相关函数
func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdName: pname, ProdID: id}
}
//降级方法相关函数，商品列表降级方法
func defaultProds(rsp interface{}){
	models := make([]*Services.ProdModel,0)
	var i int32
	for i = 0;i < 2;i ++{
		models = append(models,newProd(10 + i,"prodname" + strconv.Itoa(10 + int(i))))
	}
	result := rsp.(*Services.ProdListResponse)
	//result := &Services.ProdListResponse{}
	result.Data = models
	//return res,nil
}

func defaultData(rsp interface{}){
	switch t:=rsp.(type) {
	case *Services.ProdListResponse:
		defaultProds(rsp)
	case *Services.ProdDetailResponse:
		t.Data = newProd(10,"降级商品")
	}
}

type ProdsWrapper struct {
	client.Client
}

//重写call方法
func(this *ProdsWrapper)Call(ctx context.Context, req client.Request, rsp interface{},pots ...client.CallOption) error{
	cmdName := req.Service() + "." + req.Endpoint()
	configA := hystrix.CommandConfig{
		Timeout: 1000,
	}
	hystrix.ConfigureCommand(cmdName,configA)
	//Endpoint是方法名
	return hystrix.Do(cmdName,func() error{
		return this.Client.Call(ctx,req,rsp)
	}, func(e error) error {
		//defaultProds(rsp)
		fmt.Println(e)
		defaultData(rsp)
		return nil
	})
	//这里的降级方法其实i就是改rsp
}

func NewProdsWrapper(c client.Client) client.Client{
	return &ProdsWrapper{c}
}
