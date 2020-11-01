package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"miniserver/ProdService"
	"miniserver/model"
)

/*
商品API
*/

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("localhost:8500"), //consul运行的地址
	)
	ginRouter := gin.Default()
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", func(context *gin.Context) {
			var temp model.ProdsRequest1
			err := context.Bind(&temp)
			if err != nil {
				fmt.Println(err)
				temp = model.ProdsRequest1{
					Size: 4,
				}
			}
			fmt.Println(temp)
			context.JSON(200, gin.H{
				"data": ProdService.NewProdList(temp.Size),
				"temp": temp,
			})
		})
	}
	server := web.NewService(
		web.Name("myConsul"), //服务名
		web.Address(":8002"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)
	//server.Init()
	server.Run()
}
