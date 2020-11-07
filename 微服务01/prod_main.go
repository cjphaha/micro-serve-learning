package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"miniserver/Services"
	"miniserver/Weblib"
)

/*
商品API
*/

type logWrapper struct {
	client.Client
}
//重写call方法
func(o *logWrapper)Call(ctx context.Context, req client.Request, rsp interface{},pots ...client.CallOption) error{
	fmt.Println("调用接口")
	o.Client.Call(ctx,req,rsp)
	return nil
}

func NewLogWrapper(c client.Client) client.Client{
	return &logWrapper{c}
}
func main() {
	consulReg := consul.NewRegistry(//初始化consul
		registry.Addrs("localhost:8500"), //consul运行的地址
	)
	//商品服务,这里就是rpc相关的了
	myService := micro.NewService(
		micro.Name("Prodservew2134143.client"),
		micro.Address(":8012"),
		micro.Registry(consulReg),//注册到consul
		micro.WrapClient(NewLogWrapper),//注册装饰器
		)
	prodService := Services.NewProdService("Prodserve",myService.Client())
	httpServer := web.NewService(//对应于http，这是在consul里面注册的过程
		web.Name("httpProdService"), //服务名
		web.Address(":8011"),
		web.Handler(Weblib.NewGinRouter(prodService)),
		web.Metadata(map[string]string{"protocol" : "http"}),
		web.Registry(consulReg),
	)
	httpServer.Init()
	httpServer.Run()
}
