package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"micro-serve-2/ServiceImpl"
	"micro-serve-2/Services"
)

func main(){

	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
		)
	prodService := micro.NewService(
		micro.Name("Prodserve"),//consul名字
		micro.Address(":8010"),
		micro.Registry(consulReg),//注册到consul
	)
	prodService.Init()
	Services.RegisterProdServiceHandler(prodService.Server(),new(ServiceImpl.ProdService))//第一个参数是服务，第二个是实现类
	prodService.Run()
}