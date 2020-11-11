package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"micro-toolbox/Services"
	"micro-toolbox/ServicesImpl"
)

func main(){
	consulReg := consul.NewRegistry(registry.Addrs("localhost:8500"))
	myServices := micro.NewService(
		micro.Name("my micro service"),
		micro.Address(":8013"),
		micro.Registry(consulReg),
		)
	//将自定义的myservices注册进去
	Services.RegisterTestServiceHandler(myServices.Server(),new(ServicesImpl.TestService))
	myServices.Run()
}
