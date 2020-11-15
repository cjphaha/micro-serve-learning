package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	_ "micro-toolbox/Database"
	"micro-toolbox/Services"
	"micro-toolbox/ServicesImpl"
)

func main(){
	//2379
	//consulReg := consul.NewRegistry(registry.Addrs("localhost:8500"))
	etcdReg := etcd.NewRegistry(registry.Addrs("localhost:2379"))
	myServices := micro.NewService(
		micro.Name("api.cjp.com.service"),
		micro.Address(":8013"),
		micro.Registry(etcdReg),
		)
	//将自定义的myservices注册进去
	Services.RegisterTestServiceHandler(myServices.Server(),new(ServicesImpl.TestService))
	Services.RegisterUserServiceHandler(myServices.Server(),new(ServicesImpl.UserService))
	myServices.Run()
}
