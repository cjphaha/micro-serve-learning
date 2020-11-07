package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"miniserver/Services"
	"miniserver/Weblib"
)

/*
商品API
*/

func main() {
	consulReg := consul.NewRegistry(//初始化consul
		registry.Addrs("localhost:8500"), //consul运行的地址
	)

	//商品服务,这里就是rpc相关的了

	myService := micro.NewService(
		micro.Name("Prodservew2134143.client"),
		micro.Address(":8012"),
		micro.Registry(consulReg),//注册到consul
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
