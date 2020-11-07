package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"miniserver/Services"
)

/*
商品API
*/

func main() {
	consulReg := consul.NewRegistry(//初始化consul
		registry.Addrs("localhost:8500"), //consul运行的地址
	)
	ginRouter := gin.Default()
	//商品服务,这里就是rpc相关的了
	httpServer := web.NewService(//对应于http，这是在consul里面注册的过程
		web.Name("httpProdService"), //服务名
		web.Address(":8011"),
		web.Handler(ginRouter),
		web.Metadata(map[string]string{"protocol" : "http"}),
		web.Registry(consulReg),
	)
	myService := micro.NewService(
		micro.Name("Prodservew2134143.client"),
		micro.Address(":8012"),
		micro.Registry(consulReg),//注册到consul
		)
	//myService.Init()
	//myService.Run()
	prodService := Services.NewProdService("Prodserve",myService.Client())
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", func(c *gin.Context) {
			var prodreq Services.ProdsRequest
			err := c.Bind(&prodreq)
			fmt.Println(prodreq)
			if err != nil{//如果出错
				c.JSON(500,gin.H{
					"status":err.Error(),
				})
			}else {
				prodRes,err := prodService.GetProdsList(context.Background(),&prodreq)
				c.JSON(200,gin.H{
					"data":prodRes,
					"err":err,
				})
			}
		})
	}
	httpServer.Init()
	httpServer.Run()
}
